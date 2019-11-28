package main

 import "fmt"

 //function to print hello
 func printHello(ch chan int ) {
   fmt.Println("Hello from printHello")
   ch <- 2
 }

 func main() {

   ch  := make(chan int)
   //inline goroutine. Define a function inline and then call it.
   go func(){fmt.Println("Hello inline")
     ch <- 1
    }()
   //call a function as goroutine
   go printHello(ch)

   fmt.Println("Hello from main")
   <-ch  // wait for the processes
   <-ch  // wait for the processes
   fmt.Println("chans received!")
 }
