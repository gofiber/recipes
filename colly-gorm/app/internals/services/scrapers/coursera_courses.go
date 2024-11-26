package scrapers

import (
	"log"
	"strings"

	"fiber-colly-gorm/internals/services/database"

	"github.com/gocolly/colly"
)

// CourseraCourses is based on Colly's official examples:
// https://github.com/gocolly/colly/blob/master/_examples/coursera_courses/coursera_courses.go
func CourseraCourses() {
	c := colly.NewCollector(
		colly.AllowedDomains("coursera.org", "www.coursera.org"),

		// If Cache responses to prevent multiple download of pages
		// even if the collector is restarted uncomment this line:
		// colly.CacheDir("./coursera_cache"),
	)

	detailCollector := c.Clone()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if e.Attr("class") == "Button_1qxkboh-o_O-primary_cv02ee-o_O-md_28awn8-o_O-primaryLink_109aggg" {
			return
		}
		link := e.Attr("href")
		if !strings.HasPrefix(link, "/browse") || strings.Index(link, "=signup") > -1 || strings.Index(link, "=login") > -1 {
			return
		}
		err := e.Request.Visit(link)
		if err != nil {
			return
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	c.OnHTML(`a[data-click-key]`, func(e *colly.HTMLElement) {
		courseURL := e.Request.AbsoluteURL(e.Attr("href"))
		if strings.Index(courseURL, "coursera.org/learn") != -1 {
			err := detailCollector.Visit(courseURL)
			if err != nil {
				return
			}
		}
	})

	detailCollector.OnHTML(`#rendered-content`, func(e *colly.HTMLElement) {
		log.Println("Course found", e.Request.URL)
		title := e.ChildText("h1[data-e2e]")
		if title == "" {
			log.Println("No title found", e.Request.URL)
		}

		course := database.Course{
			Title:       title,
			URL:         e.Request.URL.String(),
			Description: e.ChildText("div.content"),
			Creator:     e.ChildText(`a[data-click-key="unified_description_page.consumer_course_page.click.hero_instructor"] > span`),
			Rating:      e.ChildText("span.css-bbd009:nth-child(2)"),
		}
		err := database.DB.Db.Create(&course).Error
		if err != nil {
			log.Println(err)
			return
		}
	})

	err := c.Visit("https://www.coursera.org/browse")
	if err != nil {
		return
	}
}
