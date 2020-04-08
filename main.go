package main

import (
	"log"
	"net/http"
	"os"
	
	"github.com/mbraeutig/go-gopher-heroku/api"
)

func main() {
	http.HandleFunc("/", api.Index)
	http.HandleFunc("/gopher", api.Gopher)
	http.HandleFunc("/healthz", api.Health)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
