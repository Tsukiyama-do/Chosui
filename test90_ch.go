package main

import (
    "fmt"
    "os"
    "bufio"
//    "strings"
)

// 文字列を1行入力
func StrStdin() (chan string) {
    fmt.Printf("Could you put your name?:")
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    a := scanner.Text()
    j := make(chan string)
    j <- a
    return j
}

func main() {
//    p := StrStdin()
//    fmt.Println(p)
    ch00 := make(chan string)

    go func() {
//      fmt.Println("This is start.")

      fmt.Println("Message received. %s\n", <-ch00)
      fmt.Println("This is end.")

    }()

    ch00 = StrStdin()

}
