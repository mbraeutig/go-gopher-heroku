package api

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/mbraeutig/go-gopher-heroku/pictures"
)

func Gopher(w http.ResponseWriter, r *http.Request) {
	// Read the gopher image file.
	f, err := os.Open("./pictures/gophercolor.png")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	// Write the gopher image to the response writer.
	if _, err := io.Copy(w, pictures.Test); err != nil {
		http.Error(w, fmt.Sprintf("Error writing response: %v", err), http.StatusInternalServerError)
	}
	w.Header().Add("Content-Type", "image/png")
}
