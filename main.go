package main

import (
	"fmt"
    // import Colly
    "github.com/gocolly/colly"
)

func main(){
	collector := colly.NewCollector()

	// whenever the collector is about to make a new request
	collector.OnRequest(func(r *colly.Request) {
		// print the url of that request
		fmt.Println("Visiting", r.URL)
	})
	collector.OnResponse(func(r *colly.Response) {
			fmt.Println("Got a response from", r.Request.URL)
	})
	collector.OnError(func(r *colly.Response, e error) {
			fmt.Println("Blimey, an error occurred!:", e)
	})

	collector.Visit("http://go-colly.org/")

}