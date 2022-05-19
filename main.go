package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	start := time.Now()

	baseURL := "www.hoodoo.digital"
	startingURL := "https://" + baseURL
	url := []string{baseURL}

	c := colly.NewCollector(
		colly.AllowedDomains(url...),
		colly.MaxDepth(0),
		colly.IgnoreRobotsTxt(),
		// colly.Debugger(colly.Debugger(&debug.LogDebugger{})),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
		// c.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	fmt.Println("Starting crawl at: ", startingURL)

	if err := c.Visit(startingURL); err != nil {
		fmt.Println("Error on start of crawl: ", err)
	}

	duration := time.Since(start)
	fmt.Println("Execution Time: ", duration)
}
