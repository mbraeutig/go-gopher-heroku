package main

import (
	"log"
	"net/http"
	"os"
)

const index = "" +
	`
<html>
    <h1>Testing Heroku</h1>
</html>
`

func main() {
	http.HandleFunc("/", Index)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, index)
}
