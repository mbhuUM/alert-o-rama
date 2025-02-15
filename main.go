package main

import (
	"log" 
	"fmt"
	"github.com/playwright-community/playwright-go"
	"time"

)
type Product struct {
    Name  string
    Price string
    URL   string
    Image string
}

func pokemoncenter() {
	pw, err := playwright.Run()
    if err != nil {
        log.Fatalf("could not start playwright: %v", err)
    }
    defer pw.Stop()

	headless := false
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
        Headless: &headless,
        Args: []string{
            fmt.Sprintf("--user-agent=%s", userAgent),
        },
    })
	if err != nil {
        log.Fatalf("could not launch browser: %v", err)
    }
    defer browser.Close()

    page, err := browser.NewPage()
    if err != nil {
        log.Fatalf("could not create page: %v", err)
    }
	 // Add headers
	 _ = page.SetExtraHTTPHeaders(map[string]string{
        "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
        "Accept-Language": "en-US,en;q=0.5",
        "Accept-Encoding": "gzip, deflate, br",
        "Connection": "keep-alive",
        "Upgrade-Insecure-Requests": "1",
        "Sec-Fetch-Dest": "document",
        "Sec-Fetch-Mode": "navigate",
        "Sec-Fetch-Site": "none",
        "Sec-Fetch-User": "?1",
    })

    if err != nil {
        log.Fatalf("could not set headers: %v", err)
    }

    if _, err = page.Goto("https://www.pokemoncenter.com/category/tcg-cards", playwright.PageGotoOptions{
        WaitUntil: playwright.WaitUntilStateNetworkidle,
        Timeout:   playwright.Float(30000),
    }); err != nil {
        log.Fatalf("could not goto: %v", err)
    }

	// Get and print the page title
	title, err := page.Title()
	if err != nil {
		log.Fatalf("could not get page title: %v", err)
	}
	fmt.Printf("Page title: %s\n", title)

	// Get and print the current URL (to check for redirects)
	url := page.URL()
	fmt.Printf("Current URL: %s\n", url)

	// Print the page content to see what we're actually getting
	content, err := page.Content()
	if err != nil {
		log.Fatalf("could not get page content: %v", err)
	}
	defer browser.Close()

	time.Sleep(1000 * time.Second)
    fmt.Println("Browser launched successfully")
	fmt.Printf("First 500 characters of page content:\n%s\n", content[:500])
}

func target() {
	pw, err := playwright.Run()
    if err != nil {
        log.Fatalf("could not start playwright: %v", err)
    }
    defer pw.Stop()

	headless := false
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
        Headless: &headless,
        Args: []string{
            fmt.Sprintf("--user-agent=%s", userAgent),
        },
    })
	if err != nil {
        log.Fatalf("could not launch browser: %v", err)
    }
    defer browser.Close()

    page, err := browser.NewPage()
    if err != nil {
        log.Fatalf("could not create page: %v", err)
    }
	 // Add headers
	 _ = page.SetExtraHTTPHeaders(map[string]string{
        "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
        "Accept-Language": "en-US,en;q=0.5",
        "Accept-Encoding": "gzip, deflate, br",
        "Connection": "keep-alive",
        "Upgrade-Insecure-Requests": "1",
        "Sec-Fetch-Dest": "document",
        "Sec-Fetch-Mode": "navigate",
        "Sec-Fetch-Site": "none",
        "Sec-Fetch-User": "?1",
    })

    if err != nil {
        log.Fatalf("could not set headers: %v", err)
    }

    // if _, err = page.Goto("https://www.target.com/p/2025-pokemon-scarlet-violet-s8-5-prismatic-evolutions-surprise-box/-/A-94336414?preselect=94336414", playwright.PageGotoOptions{
	if _, err = page.Goto("https://www.target.com/p/pok-233-mon-trading-card-game-scarlet-38-violet-8212-temporal-forces-elite-trainer-box-walking-wake/-/A-91255505", playwright.PageGotoOptions{
        WaitUntil: playwright.WaitUntilStateNetworkidle,
        Timeout:   playwright.Float(30000),
    }); err != nil {
        log.Fatalf("could not goto: %v", err)
    }

	// Get and print the page title
	title, err := page.Title()
	if err != nil {
		log.Fatalf("could not get page title: %v", err)
	}
	fmt.Printf("Page title: %s\n", title)

	// Get and print the current URL (to check for redirects)
	url := page.URL()
	fmt.Printf("Current URL: %s\n", url)

	// Print the page content to see what we're actually getting
	content, err := page.Content()
	if err != nil {
		log.Fatalf("could not get page content: %v", err)
	}
	defer browser.Close()

	time.Sleep(1000 * time.Second)
    fmt.Println("Browser launched successfully")
	fmt.Printf("First 500 characters of page content:\n%s\n", content[:500])
}


func main() {
	var option int
	fmt.Println("Select service:\n 1. Pokemon Center \n 2. Target")
	fmt.Scan(&option)

	if(option == 1) {
		pokemoncenter()
	} else if option == 2 {
		target()
	} 

}