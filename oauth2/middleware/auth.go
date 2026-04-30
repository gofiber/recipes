package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"oauth2/models"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/client"
)

// OAUTHBegin generates a CSRF state token, stores it in the session, and
// redirects the user to GitHub's authorization page.
func OAUTHBegin(ctx fiber.Ctx) error {
	models.SYSLOG.Debug("entering OAUTHBegin")

	sess, err := models.MySessionStore.Get(ctx)
	if err != nil {
		models.SYSLOG.Error("session error in OAUTHBegin", "err", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Generate a random CSRF state token
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	state := hex.EncodeToString(b)
	sess.Set("oauth-state", state)
	if err := sess.Save(); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	authURL := fmt.Sprintf(
		"https://github.com/login/oauth/authorize?client_id=%s&state=%s",
		models.ClientID, state,
	)
	return ctx.Redirect().Status(fiber.StatusTemporaryRedirect).To(authURL)
}

// OAUTHRedirect performs the GitHub OAUTH2 login sequence and stores the token in a session variable.
func OAUTHRedirect(ctx fiber.Ctx) error {
	models.SYSLOG.Debug("entering OAUTHRedirect", "url", ctx.OriginalURL())
	defer models.SYSLOG.Debug("exiting OAUTHRedirect")

	// Validate CSRF state parameter
	stateParam := ctx.Query("state", "")
	sess, err := models.MySessionStore.Get(ctx)
	if err != nil {
		models.SYSLOG.Error("session error in OAUTHRedirect", "err", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	storedState, _ := sess.Get("oauth-state").(string)
	if stateParam == "" || stateParam != storedState {
		models.SYSLOG.Warn("CSRF state mismatch", "got", stateParam, "want", storedState)
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "invalid state parameter"})
	}
	// Clear state after use
	sess.Delete("oauth-state")

	// Get the authorization code from the query param
	code := ctx.Query("code", "")
	if len(code) < 1 {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	// Exchange the code for an access token via the GitHub OAuth endpoint
	a := client.New()
	req := a.R()
	req.SetMethod("POST")
	req.SetURL(fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		models.ClientID, models.ClientSecret, code,
	))
	req.SetHeader("accept", "application/json")

	resp, clientErr := req.Send()
	if clientErr != nil {
		models.SYSLOG.Error("could not send HTTP request to GitHub", "err", clientErr)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": clientErr.Error()})
	}

	var t *models.OAuthAccessResponse
	if clientErr = resp.JSON(&t); clientErr != nil {
		models.SYSLOG.Error("could not decode GitHub response", "err", clientErr)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": clientErr.Error()})
	}

	models.SYSLOG.Debug("received GitHub token response", "status", resp.StatusCode())

	// Store token data in the session (reuse the session fetched above)
	sess.Set("oauth-scope", t.Scope)
	sess.Set("oauth-token-type", t.TokenType)
	sess.Set("oauth-token", t.AccessToken)
	if err := sess.Save(); err != nil {
		models.SYSLOG.Error("session save error", "err", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	models.SYSLOG.Debug("redirecting to /welcome.html")
	return ctx.Redirect().Status(fiber.StatusFound).To("/welcome.html")
}

// OAUTHProtected processes access attempts; if the session stored token is NULL then it sends to start page.
func OAUTHProtected(c fiber.Ctx) error {
	models.SYSLOG.Debug("entering OAUTHProtected", "url", c.OriginalURL())
	defer models.SYSLOG.Debug("exiting OAUTHProtected")

	sessData, err := models.MySessionStore.Get(c)
	if err != nil {
		models.SYSLOG.Error("session exception", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	models.SYSLOG.Debug("session fresh?", "fresh", sessData.Fresh())

	tk := sessData.Get("oauth-token")
	models.SYSLOG.Debug("session oauth-token", "value", tk)

	if tk == nil {
		sessData.Destroy()
		models.SYSLOG.Debug("token is nil, redirecting to /index.html")
		return c.Redirect().Status(fiber.StatusTemporaryRedirect).To("/index.html")
	}

	return c.Next()
}

// OAUTHGETHandler displays a "secure" page.
func OAUTHGETHandler(c fiber.Ctx) error {
	models.SYSLOG.Debug("entering OAUTHGETHandler")
	defer models.SYSLOG.Debug("exiting OAUTHGETHandler")
	return c.Render("protected", fiber.Map{})
}

// OAUTHDisconnect performs disconnection - session is destroyed.
func OAUTHDisconnect(c fiber.Ctx) error {
	models.SYSLOG.Debug("entering OAUTHDisconnect", "url", c.OriginalURL())
	defer models.SYSLOG.Debug("exiting OAUTHDisconnect")

	sessData, err := models.MySessionStore.Get(c)
	if err != nil {
		models.SYSLOG.Error("session exception", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	models.SYSLOG.Debug("destroying session")
	sessData.Destroy()

	return c.Redirect().Status(fiber.StatusTemporaryRedirect).To("/index.html")
}
