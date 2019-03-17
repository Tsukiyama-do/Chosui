package main

import (
    "fmt"
    "time"
  "golang.org/x/net/context"

)

func main() {
  ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

  var i_ok int = 0
//    var weather string    // shared memory


    con := make(chan int,5)
    var ends chan string

    go func() {   con <- 100    }()
    go func() {   ends <- "hi"    }()


    for {
      if i_ok == 0 {
        select {
        case v1 := <-con :
          fmt.Println(v1)
        case v2 := <-ends :
          fmt.Println("Hoho",v2)
        case <-ctx.Done() :
          i_ok = 1
        }

    } else {
      break
    }
    }

//    time.Sleep(1000*time.Millisecond)

/*
  con <- 5 + 7 + 8
  fmt.Println(<-con)
  fmt.Println(<-con)
  fmt.Println(<-con)
*/

}
