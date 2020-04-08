package pictures

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var Gopher bytes.Buffer
var Test *bytes.Reader

func init() {

	b, err := ioutil.ReadFile("./pictures/gophercolor.png")
	if err != nil {
		log.Fatal(err)
	}
	Test = bytes.NewReader(b)

	f, err := os.Open("./pictures/gophercolor.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := io.Copy(&Gopher, f); err != nil {
		log.Fatal(err)
	}
}
