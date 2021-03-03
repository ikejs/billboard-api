package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gocolly/colly"
)

type Song struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Rank   string `json:"rank"`
}

func scrapeSongs() []Song {
	// Instantiate default collector
	c := colly.NewCollector()
	allSongs := []Song{}

	c.OnHTML(".chart-list__element", func(e *colly.HTMLElement) {
		// Find song using an attribute selector
		// Matche elements by class
		title := e.ChildText(".chart-element__information__song")
		artist := e.ChildText(".chart-element__information__artist")
		rank := e.ChildText(".chart-element__rank__number")

		// add songs to list
		allSongs = append(allSongs, Song{
			Title:  title,
			Artist: artist,
			Rank:   rank,
		})
	})

	c.Visit("https://www.billboard.com/charts/hot-100")
	return allSongs
}

var songs []Song

func renderSongsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(songs)
}

func createFile(data []Song) {
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("songs.json", file, 0644)
}

func startServer() {
	songs = scrapeSongs()
	createFile(songs)
	http.HandleFunc("/", renderSongsHandler)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	// API
	startServer()
}
