package pictures

import (
	"bytes"
	"io"
	"log"
	"os"
)

var Gopher bytes.Buffer

func init() {
	f, err := os.Open("./pictures/gophercolor.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := io.Copy(&Gopher, f); err != nil {
		log.Fatal(err)
	}
}
