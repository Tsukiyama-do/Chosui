package main

import (

  "fmt"
  "time"
  "math/rand"
)

func main() {

  fmt.Println(rand.Intn(1e4))
  fmt.Println(time.Duration(1e6))
}
