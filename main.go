package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"encoding/csv"
    "os"
	"log" 
	"net/http" 
    "crypto/tls"  
	"github.com/playwright-community/playwright-go"

)

type Product struct {
	    Url, Image, Name, Price string
	}

func main() {
	var products []Product 

    transport := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
	c := colly.NewCollector(
        colly.AllowedDomains("www.pokemoncenter.com"),
	)

    c.WithTransport(transport)

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// called before an HTTP request is triggered
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
        r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
        r.Headers.Set("Accept-Language", "en-US,en;q=0.5")
        r.Headers.Set("Accept-Encoding", "gzip, deflate, br")
        r.Headers.Set("Connection", "keep-alive")
        r.Headers.Set("Upgrade-Insecure-Requests", "1")
        r.Headers.Set("Sec-Fetch-Dest", "document")
        r.Headers.Set("Sec-Fetch-Mode", "navigate")
        r.Headers.Set("Sec-Fetch-Site", "none")
        r.Headers.Set("Sec-Fetch-User", "?1")
        fmt.Println("Making request to:", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Response received: %s\n", r.Body)
	})

	c.OnHTML(".product--feNDW", func(e *colly.HTMLElement) {
		// initialize a new Product instance
		product := Product{}
		fmt.Println("OnHtml")
		// scrape the target data
		product.Url = e.ChildAttr("a", "href")
		product.Image = e.ChildAttr("product-image", "src")
		product.Name = e.ChildText(".product-title")
		product.Price = e.ChildText(".product-price")

		// add the product instance with scraped data to the list of products
		products = append(products, product)
	})
	
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

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36"
	
	
	c.Visit("https://www.pokemoncenter.com/category/tcg-cards")
}