package main
import (
  "fmt"
  "sync"
  "time"
  "strconv"
  "log"
)

func cleanup() {
    if r := recover(); r != nil {
      fmt.Println("Recovered in cleanup: ", r)
    }

}


func main() {
  var mu sync.Mutex

    counter := 0
    defer cleanup()
    for i := 0; i < 10; i++ {
        mu.Lock()
        go func(j int) {
          log.Println("func:", strconv.Itoa(j), time.Now().String())
          defer mu.Unlock()
            counter++
            log.Printf(" %d=%d", j, counter)
//            fmt.Print("*")
          if i == 5 {
            panic("Oh, my goodness!")
          }

        }(i)
    }


    time.Sleep(5 * time.Second)
    fmt.Printf("\n%d\n", counter)
}
