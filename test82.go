package main

import (
  "fmt"
  "sync"
//  "time"
)

func main() {
    var mu sync.Mutex
    wg := sync.WaitGroup{}
    counter := 0
    for i := 0; i < 100; i++ {
        wg.Add(1)
        mu.Lock()
        go func() {
            defer mu.Unlock()
            counter++
            fmt.Printf(" %d", counter)
            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println(counter)
}
