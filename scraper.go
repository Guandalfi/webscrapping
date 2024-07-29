package main

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly"

	"fmt"
)

type manga struct {
	url             string
	manga           string
	ultimo_capitulo float64
}

func main() {

	var capitulos []float64

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

	c.OnHTML(".chapter-name", func(e *colly.HTMLElement) {
		//e.Request.Visit(e.Attr("href"))
		//manga := manga{}

		capitulo := e.Text[8:]

		capitulo = strings.Replace(capitulo, "-", ".", -1)
		capitulo = strings.Replace(capitulo, ",", ".", -1)

		capitulo_numero, err := strconv.ParseFloat(capitulo, 64)
		if err != nil {
			// ... handle error
			panic(err)
		}

		capitulos = append(capitulos, capitulo_numero)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Acabou", r.Request.URL)
	})

	/*c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})*/

	c.Visit("https://chapmanganelo.com/manga-lu126440")
	fmt.Println(capitulos)
}
