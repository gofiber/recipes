package main

import (
	"log"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v3"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main() {
	// Create a new language bundle with the default language (English).
	bundle := i18n.NewBundle(language.English)

	// Register toml unmarshal function.
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	// Load translations for other languages.
	// No need to load active.en.toml since we are providing default translations.
	bundle.MustLoadMessageFile("./lang/active.es.toml")
	bundle.MustLoadMessageFile("./lang/active.ru.toml")

	// Create a new engine by passing the template folder
	// and template extension using <engine>.New(dir, ext string)
	engine := html.New("./templates", ".html")

	// Reload the templates on each render in development mode.
	if os.Getenv("ENV") == "development" {
		engine.Reload(true)
	}

	// After you created your engine, you can pass it to Fiber's Views Engine
	app := fiber.New(fiber.Config{Views: engine})

	// Render template.
	app.Get("/", func(c fiber.Ctx) error {
		lang := c.Query("lang")            // parse language from query
		accept := c.Get("Accept-Language") // or, parse language from header

		// Create a new localizer.
		localizer := i18n.NewLocalizer(bundle, lang, accept)

		// Set default user name (for example).
		name := c.Query("name")
		if name == "" {
			name = "Bob"
		}

		// Set title message.
		helloPerson, err := localizer.Localize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "HelloPerson",
				Other: "Hello {{.Name}}",
			},
			TemplateData: map[string]string{
				"Name": name,
			},
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Parse and set unread count of emails.
		unreadEmailCount, err := strconv.ParseInt(c.Query("unread"), 10, 64)
		if err != nil {
			unreadEmailCount = 0
		}

		// Set your own message for unread emails.
		myUnreadEmails, err := localizer.Localize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:          "MyUnreadEmails",
				Description: "The number of unread emails I have",
				One:         "I have {{.PluralCount}} unread email.",
				Other:       "I have {{.PluralCount}} unread emails.",
			},
			PluralCount: unreadEmailCount,
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Set other personal message for unread emails.
		personUnreadEmails, err := localizer.Localize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:          "PersonUnreadEmails",
				Description: "The number of unread emails a person has",
				One:         "{{.Name}} has {{.UnreadEmailCount}} unread email.",
				Other:       "{{.Name}} has {{.UnreadEmailCount}} unread emails.",
			},
			PluralCount: unreadEmailCount,
			TemplateData: map[string]interface{}{
				"Name":             name,
				"UnreadEmailCount": unreadEmailCount,
			},
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Return rendered template.
		return c.Render("index", fiber.Map{
			"Title": helloPerson,
			"Paragraphs": []string{
				myUnreadEmails,
				personUnreadEmails,
			},
		})
	})

	// Start server on port 3000.
	log.Fatal(app.Listen(":3000"))
}
