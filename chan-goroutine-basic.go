package main

import (
    "fmt"
//    "time"
)

func main() {

    var weather string    // shared memory

    con := make(chan int)
    ends := make(chan int)
//    defer close(con)
    go func() {
      defer close(con)
      <-con     //  wait until update is completed!  // communications
      fmt.Println("Today is", weather)
      ends <- 0  //  notify if main program can end itself or not.
    }()


    // q に何か入るまで待つ

    go func() {
      if weather == "" {
        weather = "Sunny"       // shared memory being updated
        con <- 0                // completed or not to update
      } else if weather == "Sunny" {
        weather = "Cloudy"      // shared memory being updated
        con <- 1                // completed or not to update
      }

    }()

    <- ends   // wait until final operation is finished. //  communications

//    time.Sleep(1000*time.Millisecond)

}
