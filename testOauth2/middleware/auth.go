package middleware

import (
	"fmt"
	"testOauth2/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// OAUTHRedirect performs the GitHub OAUTH2 login sequence and stored the token in a session variable
func OAUTHRedirect(ctx *fiber.Ctx) error {

	models.SYSLOG.Tracef("entering OAUTHRedirect; original URL: %v", ctx.OriginalURL())
	defer models.SYSLOG.Trace("exiting OAUTHRedirect")

	// First, we need to get the value of the `code` query param
	code := ctx.Query("code", "")
	if len(code) < 1 {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	// Next, lets for the HTTP request to call the github oauth enpoint	to get our access token

	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod(fiber.MethodPost)
	req.Header.Set("accept", "application/json")
	req.SetRequestURI(fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", models.ClientID, models.ClientSecret, code))
	if err := a.Parse(); err != nil {
		models.SYSLOG.Errorf("could not create HTTP request: %v", err)
	}

	var retCode int
	var retBody []byte
	var errs []error
	// Send out the HTTP request
	var t *models.OAuthAccessResponse

	if retCode, retBody, errs = a.Struct(&t); len(errs) > 0 {
		models.SYSLOG.Tracef("received: %v", string(retBody))
		models.SYSLOG.Errorf("could not send HTTP request: %v", errs)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	models.SYSLOG.Tracef("received : %v %v %v", retCode, string(retBody), errs)

	var sess *session.Session
	var err error
	// Finally, send a response to redirect the user to the "welcome" page with the access token
	if sess, err = models.MySessionStore.Get(ctx); err == nil {
		sess.Set("token", t.AccessToken)
		models.SYSLOG.Tracef("setting session token %v", t.AccessToken)
		sessData, _ := models.MySessionStore.Get(ctx)
		defer sessData.Save()
		//models.MySessionStore.RegisterType(models.OAuthAccessResponse)
		sessData.Set("oauth-scope", t.Scope)
		sessData.Set("oauth-token-type", t.TokenType)
		sessData.Set("oauth-token", t.AccessToken)

		if err != nil {
			models.SYSLOG.Errorf("session saving exception %v", err)
		}
		models.SYSLOG.Tracef("redirecting to /welcome.html?access_token=%v", t.AccessToken)
		//		return ctx.Redirect("/welcome.html?access_token="+t.AccessToken, fiber.StatusFound)
		return ctx.Redirect("/welcome.html", fiber.StatusFound)
	}

	models.SYSLOG.Tracef("redirecting to /")
	return ctx.Redirect("/", fiber.StatusTemporaryRedirect)
}

// OAUTHProtected processes access attempts; if the session stored token is NULL then it sends to start page
func OAUTHProtected(c *fiber.Ctx) error {
	models.SYSLOG.Tracef("entering OAUTHProtected; original URL: %v", c.OriginalURL())
	defer models.SYSLOG.Trace("exiting OAUTHProtected")

	sessData, err := models.MySessionStore.Get(c)
	if err != nil {
		models.SYSLOG.Errorf("session exception %v", err)
		panic(err)
	}

	// for debug purposes - inspect the session variables
	models.SYSLOG.Tracef("session id fresh ? %v", sessData.Fresh())

	models.SYSLOG.Trace("trying to get 'oauth-scope' value")
	tk := sessData.Get("oauth-scope")
	models.SYSLOG.Tracef("session stored 'oauth-scope' is %v", tk)

	models.SYSLOG.Trace("trying to get 'oauth-token-type' value")
	tk = sessData.Get("oauth-token-type")
	models.SYSLOG.Tracef("session stored 'oauth-token-type' is %v", tk)

	tk = sessData.Get("oauth-token")
	models.SYSLOG.Tracef("session stored 'oauth-token' is %v", tk)

	if tk == nil {
		sessData.Destroy()
		models.SYSLOG.Tracef("token is NULL")
		return c.Redirect("/index.html", fiber.StatusTemporaryRedirect)
	}

	return c.Next()
}

// OAUTHGETHandler displays a "secure" page
func OAUTHGETHandler(c *fiber.Ctx) error {
	models.SYSLOG.Trace("entering OAUTHGETHandler")
	defer models.SYSLOG.Trace("exiting OAUTHGETHandler")
	return c.Render("protected", fiber.Map{})
}

// OAUTHDisconnect performs disconnection - session is destroyed and
func OAUTHDisconnect(c *fiber.Ctx) error {
	models.SYSLOG.Tracef("entering OAUTHDisconnect; original URL: %v", c.OriginalURL())
	defer models.SYSLOG.Trace("exiting OAUTHDisconnect")
	sessData, err := models.MySessionStore.Get(c)
	if err != nil {
		models.SYSLOG.Errorf("session exception %v", err)
		panic(err)
	}

	// for debug purposes - inspect the session variables
	models.SYSLOG.Tracef("session id fresh ? %v", sessData.Fresh())

	models.SYSLOG.Trace("trying to get 'oauth-scope' value")
	tk := sessData.Get("oauth-scope")
	models.SYSLOG.Tracef("session stored 'oauth-scope' is %v", tk)

	models.SYSLOG.Trace("trying to get 'oauth-token-type' value")
	tk = sessData.Get("oauth-token-type")
	models.SYSLOG.Tracef("session stored 'oauth-token-type' is %v", tk)

	tk = sessData.Get("oauth-token")
	models.SYSLOG.Tracef("session stored 'oauth-token' is %v", tk)

	sessData.Destroy()

	return c.Redirect("/index.html", fiber.StatusTemporaryRedirect)
}
