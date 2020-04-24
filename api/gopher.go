package api

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var gopher []byte

func init() {
	if err := readFile("./images/gophercolor.png", &gopher); err != nil {
		log.Fatal(err)
	}
}

func readFile(filename string, data *[]byte) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	*data = b
	return nil
}

func Gopher(w http.ResponseWriter, r *http.Request) {
	// Write the gopher image to the response writer.
	if _, err := io.Copy(w, bytes.NewReader(gopher)); err != nil {
		http.Error(w, fmt.Sprintf("Error writing response: %v", err), http.StatusInternalServerError)
	}
	w.Header().Add("Content-Type", "image/png")
}
