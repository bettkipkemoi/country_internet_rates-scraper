package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Item struct {
	Country string `json:"country"`
	Cost    string `json:"cost"`
}

func main() {
	scrapeUrl := "https://worldpopulationreview.com/country-rankings/internet-cost-by-country"

	c := colly.NewCollector(colly.AllowedDomains("https://worldpopulationreview.com", "worldpopulationreview.com"))
	items := []Item{}

	c.OnHTML("h1.text-2xl", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	c.OnHTML("div.mb-5", func(h *colly.HTMLElement) {
		selection := h.DOM
		i := Item{
			Country: selection.Find("a[href]").Text(),
			Cost:    h.ChildText("td"),
		}
		items = append(items, i)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(fmt.Printf("Scraping %s", r.URL))
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("There was an error while scraping: %s\n", err.Error())
	})
	c.Visit(scrapeUrl)
	fmt.Println(items)

}
