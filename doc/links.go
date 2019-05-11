package doc

import (
	"github.com/gocolly/colly"
	"github.com/pitw/doc-parser-proffix/model"
	"log"
	"strings"
)

func GetDocLinks() (doclinks []model.DocLink) {
	// Instantiate default collector
	c := colly.NewCollector()
	// On every a element which has href attribute call callback
	c.OnHTML("a", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		fullLink := "https://www.proffix.net/Portals/0/content/REST%20API/export/" + link

		// Print link
		if strings.HasSuffix(link, "html") {
			doclinks = append(doclinks, model.DocLink{
				Name: e.Text,
				Link: fullLink,
			})

		}

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Printf("Visiting %v", r.URL.String())
	})

	c.Visit("https://www.proffix.net/Portals/0/content/REST%20API/export/proffix_rest_api_entwicklerhandbuch_content.html")

	return doclinks
}