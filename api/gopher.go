package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/mbraeutig/go-gopher-heroku/pictures"
)

func Gopher(w http.ResponseWriter, r *http.Request) {
	// Write the gopher image to the response writer.
	if _, err := io.Copy(w, bytes.NewReader(pictures.Gopher)); err != nil {
		http.Error(w, fmt.Sprintf("Error writing response: %v", err), http.StatusInternalServerError)
	}
	w.Header().Add("Content-Type", "image/png")
}
