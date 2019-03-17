package main
import (
  "fmt"
  "time"
)

func main(){

  var msg  string = "Hello"
  go func(msg string){
    fmt.Println(msg)
  }()

  msg = "Good Morning"
  time.Sleep(100 * time.Millisecond)

}
