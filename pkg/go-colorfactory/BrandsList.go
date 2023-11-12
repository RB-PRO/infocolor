package gocolorfactory

import (
	"fmt"

	"github.com/gocolly/colly"
)

func Brands() (brands []string, Err error) {

	// Create a collector
	c := colly.NewCollector()

	// Set HTML callback
	// Won't be called if error occurs
	c.OnHTML(`li[class="flex xs6 sm4 md3 d-flex"]>a`, func(e *colly.HTMLElement) {
		brand, _ := e.DOM.Attr("href")
		if brand != "" {
			brands = append(brands, brand)
		}
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		Err = fmt.Errorf("request URL: %v failed with response: %v\nError: %v", r.Request.URL, r, err)
	})

	// Start scraping
	c.Visit(URL + "/avtokraski/")
	return brands, Err
}
