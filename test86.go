package main
import (

  "log"
  "sync"
  "time"
  "strconv"

)

var wg sync.WaitGroup

func cleanup() {

  if r := recover() ;  r != nil {
    log.Println("Recovered, ", r)
  }
}

func say(s string) {
  defer wg.Done()
  defer cleanup()
  for i := 0 ; i < 10 ; i++ {
    time.Sleep(100*time.Millisecond)
    log.Println(s)
    log.Println(strconv.Itoa(i), " ", time.Now().String())

    if i == 5 {
      panic("Oh my")
    }
  }
}

func main() {
  wg.Add(1)
  go say("Hey")
  wg.Add(1)
  go say("Wait")
  wg.Wait()
}
