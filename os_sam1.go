package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewReader(os.Stdin)
	fmt.Printf("Hello, %s", w)
	fmt.Printf("World! %s", w)
	//	w.Flush() // Don't forget to flush!

}
