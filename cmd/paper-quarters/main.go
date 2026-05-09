package main

import (
	"log"

	"paper_quarters/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
