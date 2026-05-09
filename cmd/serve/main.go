package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", "localhost:8080", "HTTP server address")
	dir := flag.String("dir", "web", "directory to serve")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*dir)))

	fmt.Printf("Serving %s at http://%s/\n", *dir, *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
