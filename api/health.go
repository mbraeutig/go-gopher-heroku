package api

import (
	"io"
	"log"
	"net/http"
)

// Health - Health Check
func Health(w http.ResponseWriter, r *http.Request) {
	log.Print("API Health is OK")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}
