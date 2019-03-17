package main

import (
	"fmt"
	"strings"
  "time"
)

func Job(input chan string) (output chan string) {
	output = make(chan string)
	go func() {
		defer close(output)
		for in := range input {
			output <- strings.ToUpper(in)
		}
	}()
	return output
}

func main() {
	input := make(chan string)
	go func() {
		defer close(input)
    time.Sleep(1*time.Second)
		input <- "hoge"
		input <- "moge"
	}()
	for out := range Job(input) {
		if len(out) != 0 {
			fmt.Println(out)
		}
	}
}
