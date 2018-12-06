package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("td.title", func(e *colly.HTMLElement) {
		e.ForEach("a.storylink", func(_ int, elem *colly.HTMLElement) {
			fmt.Printf("%s\n", elem.Text)
		})
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("failed with response:", r, "\nError:", err)
	})

	c.Visit("https://news.ycombinator.com/news?p=1")
}
