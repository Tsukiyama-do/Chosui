package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

func main() {
	//	r := strings.NewReader("some io.Reader stream to be read\n")
	r := bytes.NewBufferStrings("some io.Reader stream to be read\n")
	var buf bytes.Buffer
	tee := io.TeeReader(r, &buf)

	printall := func(r io.Reader) {
		b, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", b)
	}

	printall(tee)
	printall(&buf)

}
