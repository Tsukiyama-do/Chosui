package main

import (
    "fmt"
    "strings"
    "time"
)

func main() {
    q := make(chan string, 2)
    output := make(chan string,2)
    go func() {
      defer close(q)
      for in := range q {
        output <- strings.ToUpper(in)
      }
    }()

    // q に何か入るまで待つ
    q <- "hello"
    q <- "hi"

    fmt.Println(<-output)
    fmt.Println(<-output)

    close(q)

    time.Sleep(1*time.Second)
    q <- "boy"
    q <- "girl"

    fmt.Println(<-output)
    fmt.Println(<-output)



}
