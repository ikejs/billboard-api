package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type song struct {
	title  string
	artist string
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector()
	songs := []song{}

	c.OnHTML(".chart-element__information", func(e *colly.HTMLElement) {
		// Find song using an attribute selector
		// Matche elements by class
		title := e.ChildText(".chart-element__information")
		artist := e.ChildText(".text--truncate")

		// Print song info
		// fmt.Printf("Song found: %q -> %s\n", title, artist)
		songs = append(songs, song{title, artist})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		fmt.Println(songs)
	})

	c.Visit("https://www.billboard.com/charts/hot-100")
}
