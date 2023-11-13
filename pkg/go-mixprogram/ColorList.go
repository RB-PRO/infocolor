package gomixprogram

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

const URL string = "https://mixprogram.ru/car_brands/%s/view-fcolors?page=%d"

func Colors(brand string, pagecount int) (colors []string, Err error) {

	// Create a collector
	c := colly.NewCollector()

	// Set HTML callback // Won't be called if error occurs
	c.OnHTML(`div[class="field ft_parent f_parent_colors_id"]>div>a`, func(e *colly.HTMLElement) {
		text := e.DOM.Text()
		text = strings.Replace(text, "GAZ", "", 1)
		text = strings.Replace(text, "MINI", "", 1)
		text = strings.Replace(text, "MINI", "", 1)
		text = strings.Replace(text, "DAEWOO", "", 1)
		text = strings.Replace(text, "UAZ", "", 1)
		text = strings.Replace(text, "GREATWALL", "", 1)
		text = strings.Replace(text, "TAGAZ", "", 1)
		text = strings.Replace(text, "LIFAN", "", 1)
		text = strings.Replace(text, "|", "", 1)
		text = strings.Replace(text, ":", "", 1)
		text = strings.Replace(text, "/", "", 1)
		text = strings.TrimSpace(text)

		colors = append(colors, text)
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		Err = fmt.Errorf("request URL: %v failed with response: %v\nError: %v",
			r.Request.URL, r, err)
	})

	for page := 1; page <= pagecount; page++ { // Start scraping
		url := fmt.Sprintf(URL, brand, page)
		fmt.Println(url)
		c.Visit(url)
	}
	return colors, Err
}
