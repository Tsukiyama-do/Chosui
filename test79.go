package main
import (
  "fmt"
//  "time"
  "sync"
)

var wg = sync.WaitGroup{}

func main(){

  var msg  string = "Hello"
  wg.Add(1)
  go func(msg string){
    fmt.Println(msg)
    wg.Done()
  }(msg)
  msg = "Good Morning"

  go func(msg string){
    fmt.Println(msg)
    wg.Done()
  }(msg)

  wg.Wait()

}
