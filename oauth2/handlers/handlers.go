package handlers

import (
	"oauth2/models"

	"github.com/gofiber/fiber/v3"
)

// HTMLPages will render and return "public" pages
func HTMLPages(c fiber.Ctx) error {
	models.SYSLOG.Tracef("entering HtmlPages; original URL: %v", c.OriginalURL())
	defer models.SYSLOG.Trace("exiting HtmlPages")

	// models.SYSLOG.Tracef("inspecting header: %v", &c.Request().Header)

	p := c.Path()
	switch p {
	case "/index.html":
		return c.Render("index", fiber.Map{
			"ClientID": models.ClientID,
		})

	case "/welcome.html":
		sessData, err := models.MySessionStore.Get(c)
		if err != nil {
			return func() {
				__fiberRedirectTarget := "/errpage.html"
				__fiberRedirectStatus := fiber.StatusInternalServerError
				return c.Redirect().Status(__fiberRedirectStatus).To(__fiberRedirectTarget)
			}()
		}

		return c.Render("welcome", fiber.Map{
			"tokenValue": sessData.Get("oauth-token"),
		})

	case "/errpage.html":
		return c.Render("errpage", fiber.Map{})
	}

	return c.Next()
}
