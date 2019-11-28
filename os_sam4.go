package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	var r_f, r_s, r_b io.Reader
	var buf bytes.Buffer
	var err error

	p := make([]byte, 25)
	r_r, err := os.OpenFile("sample1.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	r_w, err := os.OpenFile("samplew.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	defer r_r.Close()
	defer r_w.Close()

	r_f = r_r
	r_s = strings.NewReader("Hirose is here. Fishing is the best hobby to communicate with fishes.")

	_, err = buf.WriteString("Masks will protect you from viruses around you.")

	r_b = &buf

	_, err = r_f.Read(p)
	log.Printf("Data part of file is %s\n", string(p))

	_, err = r_s.Read(p)
	log.Printf("Data part of String is %s\n", string(p))

	_, err = r_b.Read(p)
	log.Printf("Data part of Buffer is %s\n", string(p))

	bb, err := ioutil.ReadAll(r_s)
	log.Printf("Data part of String ioutil is %s\n", string(bb))

	_, err = io.Copy(r_w, r_b)

}
