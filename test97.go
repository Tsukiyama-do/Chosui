package main

import (

  "fmt"
  "time"
)

func makeCh() chan int {
  return make(chan int, 5)

}

func recvCh(recv <-chan int) int {
  return <-recv
}


func main() {

  ch := makeCh()
  go func(ch chan<- int ) {
    time.Sleep(800*time.Millisecond)
    ch <- 200
    } (ch)
  fmt.Println(recvCh(ch))

  ch <- 3
  ch <- 4
  ch <- 5
  close(ch)
    for v := range ch {
      fmt.Printf("Channel %d \n", v)
    }

}
