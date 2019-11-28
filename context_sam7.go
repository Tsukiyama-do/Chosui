package main

import (
  "context"
  "fmt"
  "math/rand"
  "time"
)


//Slow function
func sleepRandom(fromFunction string, i_multi int ,ch chan int,ss string, rh string) {
  //defer cleanup
  defer func() { fmt.Println(fromFunction, "sleepRandom complete due to defer ", ss, rh) }()

  //Perform a slow task
  //For illustration purpose,
  //Sleep here for random ms

  seed := time.Now().UnixNano()
  r := rand.New(rand.NewSource(seed))
  randomNumber := r.Intn(100)
  sleeptime := (randomNumber + 100 ) * i_multi
  fmt.Println(fromFunction, "Starting sleep for", sleeptime, "ms ",ss, rh)
  time.Sleep(time.Duration(sleeptime) * time.Millisecond)
  fmt.Println(fromFunction, "Waking up, slept for ", sleeptime, "ms ", ss, rh)

  //write on the channel if it was passed in
  if ch != nil {
    ch <- sleeptime
  }
}


func request_handle(ctx context.Context, ch chan bool,ss string, rh string){

  defer func() {
    fmt.Println("request_handle complete due to defer ", ss, rh)
    ch <- true
  }()

  //Make a channel
  sleeptimeChan := make(chan int)

  // Application area like waiting response from backyard.

  go sleepRandom("request_handle",5, sleeptimeChan, ss, rh)

  // Watching
  select {
  case <-ctx.Done():

    fmt.Println("request_handle cancelled. ", ss, rh)

  case sleeptime := <-sleeptimeChan:
    //This case is selected when processing finishes before the context is cancelled
    fmt.Println("Slept for ", sleeptime, "ms ", ss, rh)
  }



}

func session_handle(ctx context.Context, ss string) {

//  ctxWithTimeout1, cancelFunction1 := context.WithTimeout(ctx, time.Duration(10000)*time.Millisecond)  // grandchild context
//  ctxWithTimeout2, cancelFunction2 := context.WithTimeout(ctx, time.Duration(5000)*time.Millisecond)  // grandchild context

  ctxWithCancel1, cancelFunction1 := context.WithCancel(ctx)
  ctxWithCancel2, cancelFunction2 := context.WithCancel(ctx)

  defer func() {
    fmt.Println("session_handle ended due to defer ", ss)
    cancelFunction1()
    cancelFunction2()
  }()

  ch1 := make(chan bool)
  ch2 := make(chan bool)

  go request_handle(ctxWithCancel1, ch1, ss,"rh1")    // grand child context  - request handling
  go request_handle(ctxWithCancel2, ch2, ss,"rh2")    // grand child context


  select {
  case <-ctx.Done():    // self is done, or killed from parent.

    fmt.Println("session_handle: self is stopped from parent context. ", ss)

  case <-ch1:            // grand child context is returned faster than parent.

    fmt.Println("session_handle: ch returned, i.e. grand-child context ended faster. ", ss)
  case <-ch2:            // grand child context is returned faster than parent.

      fmt.Println("session_handle: ch returned, i.e. grand-child context ended faster. ", ss)
  }

}


func main() {


  //Make a background context
  ctx := context.Background()
  //Derive a context with cancel
  ctxWithCancel1, cancelFunction1 := context.WithCancel(ctx)
  ctxWithCancel2, cancelFunction2 := context.WithCancel(ctx)

  //defer canceling so that all the resources are freed up
  //For this and the derived contexts
  defer func() {
    fmt.Println("Main Defer: canceling context")
    cancelFunction1()
    cancelFunction2()
  }()

  //Cancel context after a random time
  //This cancels the request after a random timeout
  //If this happens, all the contexts derived from this should return
  go func() {
    sleepRandom("Main",4, nil,"main ", "main ")   // sleeper 1
    cancelFunction1()      // cancel child context
    cancelFunction2()      // cancel child context
    fmt.Println("Main Sleep complete. canceling context")
  }()
  //Do work
  go session_handle(ctxWithCancel1,"ss1")     // child context - session open
  go session_handle(ctxWithCancel2,"ss2")     // child context

  sleepRandom("Main",9, nil,"main ", "main ")   // sleeper 1

}
