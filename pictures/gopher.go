package pictures

import (
	"io/ioutil"
	"log"
)

var Gopher []byte

func init() {
	if err := readFile("./public/gophercolor.png", &Gopher); err != nil {
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
