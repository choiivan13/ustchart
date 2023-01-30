package main

import (
	"fmt"

	"github.com/choiivan13/ustchart/backend/internal/scraper"
)

func main() {
	c := scraper.NewScraper("123", 123)
	c.Scrape()

	fmt.Print(c.Courses)
}
