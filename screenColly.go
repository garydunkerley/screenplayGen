package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("imsdb.com/TV/Futurama"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		
		if strings.HasPrefix(link, "/TV Transcripts/Futurama") || strings.HasPrefix(link, /trasncripts/Futurama) {
			e.Request.Visit(link)
		}

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("imsdb.com/TV/Futurama.html")
}
