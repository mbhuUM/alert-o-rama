package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"encoding/csv"
    "os"
	"log" 
	"sync"

)

type Product struct {
    Url, Image, Name, Price string
}

func main() {
	// define a sync to filter visited URLs
	var visitedUrls sync.Map

    c := colly.NewCollector(
        colly.AllowedDomains("www.scrapingcourse.com"),)

	
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36"

    // set up the proxy
    err := c.SetProxy("http://35.185.196.38:3128")
    if err != nil {
        log.Fatal(err)
    }

	// OnError callback
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	
    // called before an HTTP request is triggered
    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting: ", r.URL)
    })

	var products []Product

	 // OnHTML callback
	 c.OnHTML("li.product", func(e *colly.HTMLElement) {

        // initialize a new Product instance
        product := Product{}

        // scrape the target data
        product.Url = e.ChildAttr("a", "href")
        product.Image = e.ChildAttr("img", "src")
        product.Name = e.ChildText(".product-name")
        product.Price = e.ChildText(".price")

        // add the product instance with scraped data to the list of products
        products = append(products, product)
    })

	  // OnHTML callback for handling pagination
	  c.OnHTML("a.next", func(e *colly.HTMLElement) {

        // extract the next page URL from the next button
        nextPage := e.Attr("href")

        // check if the nextPage URL has been visited
        if _, found := visitedUrls.Load(nextPage); !found {
            fmt.Println("scraping:", nextPage)
            // mark the URL as visited
            visitedUrls.Store(nextPage, struct{}{})
            // visit the next page
            e.Request.Visit(nextPage)
        }
    })

	//writing to a csv
	c.OnScraped(func(r *colly.Response) {

        // open the CSV file
        file, err := os.Create("products.csv")
        if err != nil {
            log.Fatalln("Failed to create output CSV file", err)
        }
        defer file.Close()

        // initialize a file writer
        writer := csv.NewWriter(file)

        // write the CSV headers
        headers := []string{
            "Url",
            "Image",
            "Name",
            "Price",
        }
        writer.Write(headers)

        // write each product as a CSV row
        for _, product := range products {
            // convert a Product to an array of strings
            record := []string{
                product.Url,
                product.Image,
                product.Name,
                product.Price,
            }

            // add a CSV record to the output file
            writer.Write(record)
        }
        defer writer.Flush()
    })

	// open the target URL
	// c.Visit("https://www.scrapingcourse.com/ecommerce")
	c.Visit("https://www.scrapingcourse.com/antibot-challenge")


}