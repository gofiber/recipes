package scrapers

import (
	"fiber-colly-gorm/internals/services/database"
	"log"

	"github.com/gocolly/colly"
)

func StartScraper() {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	quotes_url := "http://quotes.toscrape.com/"
	err := c.Visit(quotes_url)
	if err != nil {
		log.Println("Error visiting Page:", err)
	}

	c.OnHTML("div.quote", func(e *colly.HTMLElement) {

		newQuote := database.Quote{}

		newQuote.Text = e.ChildText("span.text")
		newQuote.Author = e.ChildText("small.author")

		err := database.DB.Db.Create(&newQuote).Error
		if err != nil {
			return
		}
	})
}
