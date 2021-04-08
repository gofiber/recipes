package main

import (
	"log"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
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

	// Reload the templates on each render, good for development
	engine.Reload(true) // Optional. Default: false

	// After you created your engine, you can pass it to Fiber's Views Engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Render template.
	app.Get("/", func(c *fiber.Ctx) error {
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
		helloPerson := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "HelloPerson",
				Other: "Hello {{.Name}}",
			},
			TemplateData: map[string]string{
				"Name": name,
			},
		})

		// Parse and set unread count of emails.
		unreadEmailCount, _ := strconv.ParseInt(c.Query("unread"), 10, 64)

		// Set your own message for unread emails.
		myUnreadEmails := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:          "MyUnreadEmails",
				Description: "The number of unread emails I have",
				One:         "I have {{.PluralCount}} unread email.",
				Other:       "I have {{.PluralCount}} unread emails.",
			},
			PluralCount: unreadEmailCount,
		})

		// Set other personal message for unread emails.
		personUnreadEmails := localizer.MustLocalize(&i18n.LocalizeConfig{
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
