package main

import (
	"github.com/gocolly/colly"

	"fmt"
)

func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visitando", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("DEU ERRO:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visitado ", r.Request.URL)
	})

	c.OnHTML("ul", func(e *colly.HTMLElement) {
		//e.Request.Visit(e.Attr("href"))
		link := e.Text[0:1]
		fmt.Println(link)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Acabou", r.Request.URL)
	})

	/*c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})*/

	c.Visit("https://chapmanganelo.com/manga-lu126440")
}
