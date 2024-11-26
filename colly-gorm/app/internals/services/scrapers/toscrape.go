package scrapers

import (
	"log"

	"fiber-colly-gorm/internals/services/database"

	"github.com/gocolly/colly"
)

func Quotes() {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	quotes_url := "http://quotes.toscrape.com/"
	err := c.Visit(quotes_url)
	if err != nil {
		log.Println("Error visiting Page:", err)
	}

	c.OnHTML("li.next a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnHTML("div.quote", func(e *colly.HTMLElement) {
		newQuote := database.Quote{}

		newQuote.Text = e.ChildText("span.text")
		newQuote.Author = e.ChildText("small.author")

		err := database.DB.Db.Create(&newQuote).Error
		if err != nil {
			return
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	c.Visit("https://quotes.toscrape.com/")
}
