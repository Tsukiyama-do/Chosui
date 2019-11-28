package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("Hello World (Buffer)\n")) // 溜め
	fmt.Println(buffer.String())
}
