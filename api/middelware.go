package api

import (
	"log"
	"net/http"
)

func LogFunc(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		f(w, r)
	}
}
