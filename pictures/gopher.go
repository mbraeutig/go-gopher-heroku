package pictures

import (
	"io/ioutil"
	"log"
)

var Gopher []byte

func init() {
	b, err := ioutil.ReadFile("./pictures/gophercolor.png")
	if err != nil {
		log.Fatal(err)
	}
	Gopher = b
}
