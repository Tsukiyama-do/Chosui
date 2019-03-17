package main

import (
  "fmt"
  "time"
)

func main() {
    messages := make(chan string)

    go func() {
      msg := <- messages
      fmt.Println("Answer is ", msg)

      }()
      for i := 0; i < 3 ; i++ {
        fmt.Println(" ",i+1, " second(s)")
        time.Sleep(1*time.Second)
      }

      messages <- "Hi"
//    msg := <-messages
      for i := 0; i < 3 ; i++ {
        fmt.Println(" ",i+1, " second(s)")
        time.Sleep(1*time.Second)
      }

//    fmt.Println(msg)
}
