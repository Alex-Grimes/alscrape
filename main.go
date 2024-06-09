package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Make HTTP GET request
	res, err := http.Get("https://anilist.co/search/anime/top-100")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Parse the HTML
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the anime titles
	doc.Find(".title").Each(func(i int, s *goquery.Selection) {
		// Extract the text of the title
		title := s.Text()
		// Print the title
		fmt.Println("Anime Title:", title)
	})

	// Find the anime scores
	doc.Find(".rank").Each(func(i int, s *goquery.Selection) {
		// Extract the text of the score
		score := s.Text()
		// Print the score
		fmt.Println("Score:", score)
	})

	// Find  the anime images
	doc.Find(".cover").Each(func(i int, s *goquery.Selection) {
		// Extract the image source
		imageSrc := s.Find("img").AttrOr("src", "")
		// Print the image source
		fmt.Println("Image Source:", imageSrc)
	})
}
