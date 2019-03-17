package main

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/net/context"
)

func Job(ctx context.Context, input chan string) chan string {
	output := make(chan string)
	go func() {
		defer close(output)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("cancel1:", ctx.Err())
				return
			case v, ok := <-input:
				if !ok {
					return
				}
				output <- strings.ToUpper(v)
			}
		}
	}()
	return output
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	input := make(chan string)
	go func() {
		defer close(input)
		for _, v := range []string{"hoge", "moge","hage","tsuri","otona","kodomo", "new"} {
			select {
			case <-ctx.Done():
        fmt.Println("cancel2:", ctx.Err())
				return
			case input <- v:
				//
			}
			time.Sleep(200 * time.Millisecond)
		}
		time.Sleep(100 * time.Millisecond)
	}()
	for v := range Job(ctx, input) {
		fmt.Println(v)
	}
	time.Sleep(100 * time.Millisecond)
}
