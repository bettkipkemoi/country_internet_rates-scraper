package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Item struct {
	Country     string `json:"country"`
	MonthlyCost string `json:"monthlycost"`
	CostPerMbps string `json:"costpermbps"`
}

func main() {
	//define the target url
	scrapeUrl := "https://worldpopulationreview.com/country-rankings/internet-cost-by-country"

	//instantiate a new collector object
	c := colly.NewCollector(
		colly.AllowedDomains("https://worldpopulationreview.com", "worldpopulationreview.com"))

	//initialize slice of structs to contain the scraped data
	items := []Item{}
	//scrape the header
	c.OnHTML("h1.text-2xl", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	c.OnHTML("div.mb-5", func(h *colly.HTMLElement) {

		//scrape the target data
		selection := h.DOM
		i := Item{
			Country:     selection.Find("a[href]").Text(),
			MonthlyCost: h.ChildText("td:nth-child(2)"),
			CostPerMbps: h.ChildText("td:nth-child(3)"),
		}
		//add the item instance with scraped data
		items = append(items, i)
	})

	//store the data to csv file after extraction
	c.OnScraped(func(r *colly.Response) {

		//open the csv file
		file, err := os.Create("items.csv")
		if err != nil {
			log.Fatalln("failed to create output json file", err)
		}
		defer file.Close()

		//initialize a file writer
		writer := csv.NewWriter(file)

		//write the csv headers
		headers := []string{
			"country",
			"monthlycost",
			"costpermbps",
		}

		//add a csv record to the output file
		writer.Write(headers)

		// write each product as a csv row
		for _, item := range items {
			// convert item to an array of strings
			record := []string{
				item.Country,
				item.MonthlyCost,
				item.CostPerMbps,
			}

			// add a csv record to ouput file
			writer.Write(record)
		}
		defer writer.Flush()
	})

	//a message while crawling the website
	c.OnRequest(func(r *colly.Request) {
		fmt.Println(fmt.Printf("Scraping %s", r.URL))
	})

	//print an error if not successful
	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("There was an error while scraping: %s\n", err.Error())
	})
	c.Visit(scrapeUrl)
	fmt.Println(items)
}
