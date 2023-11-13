package govseautoemali

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

// https://vseautoemali.ru/cars/chevrolet?sheet=2
// https://vseautoemali.ru/cars/%s?sheet=%d
const URL string = "https://vseautoemali.ru/cars/%s?sheet=%d"

func Colors(brand string, maxpage int) (colors []string, Err error) {

	// Create a collector
	c := colly.NewCollector()

	// Set HTML callback // Won't be called if error occurs
	c.OnHTML(`span[class="label label-code"]`, func(e *colly.HTMLElement) {
		text := e.DOM.Text()
		text = strings.TrimSpace(text)
		colors = append(colors, text)
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		Err = fmt.Errorf("request URL: %v failed with response: %v\nError: %v",
			r.Request.URL, r, err)
	})

	for page := 1; page <= maxpage; page++ { // Start scraping
		url := fmt.Sprintf(URL, brand, page)
		fmt.Println(url)
		c.Visit(url)
	}
	return colors, Err
}
