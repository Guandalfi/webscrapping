package main

import (
	"bufio"
	"log"
	"os"

	"github.com/gocolly/colly"

	"fmt"
)

type Ultimo_capitulo struct {
	Url             string
	Manga           string
	Ultimo_capitulo float64
}

type Capitulo struct {
	Manga    string
	Capitulo float64
}

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

	c.OnHTML("", func(e *colly.HTMLElement) {
		fmt.Println(e)

		/*manga := Capitulo{}

		fmt.Println(e.Attr("title"))

		capitulo := e.Text[8:]

		capitulo = strings.Replace(capitulo, "-", ".", -1)
		capitulo = strings.Replace(capitulo, ",", ".", -1)

		capitulo_numero, err := strconv.ParseFloat(capitulo, 64)
		if err != nil {
			// ... handle error
			panic(err)
		}
		manga.Manga = "teste"
		manga.Capitulo = capitulo_numero

		content, err := json.Marshal(manga)
		if err != nil {
			fmt.Println(err)
		}
		err = os.WriteFile("teste.json", content, 0644)
		if err != nil {
			log.Fatal(err)
		}*/

	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Acabou", r.Request.URL)
	})

	/*c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})*/

	//Abre arquivo com as URLs
	mangas, err := os.Open("mangas.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer mangas.Close()

	//Verifica linha a linha
	scanner := bufio.NewScanner(mangas)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		c.Visit(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
