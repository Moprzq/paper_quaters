package main

import (
	"flag"
	"log"

	"paper_quarters/internal/app"
)

func main() {
	language := flag.String("lang", "", "UI language: ru or eng")
	flag.Parse()

	if err := app.Run(*language); err != nil {
		log.Fatal(err)
	}
}
