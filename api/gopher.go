package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// Gopher prints a gopher.
func Gopher(w http.ResponseWriter, r *http.Request) {

	infos, err := ioutil.ReadDir(".")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading dir: %v", err), http.StatusInternalServerError)
		return
	}
	for _, i := range infos {
		//fmt.Print(w, i.Name())
		io.WriteString(w, "\n"+i.Name())
	}

	// Read the gopher image file.
	f, err := os.Open("gophercolor.png")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	// Write the gopher image to the response writer.
	if _, err := io.Copy(w, f); err != nil {
		http.Error(w, fmt.Sprintf("Error writing response: %v", err), http.StatusInternalServerError)
	}
	w.Header().Add("Content-Type", "image/png")
}
