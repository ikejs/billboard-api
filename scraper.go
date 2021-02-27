package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gocolly/colly"
)

type song struct {
	title  string
	artist string
}

func scrapeSongs() []song {
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
		// add songs to list
		songs = append(songs, song{title, artist})
	})

	c.Visit("https://www.billboard.com/charts/hot-100")
	return songs
}

func renderSongs(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(scrapeSongs())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func handleRequests() {
	http.HandleFunc("/", renderSongs)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {

	// API
	handleRequests()

}
