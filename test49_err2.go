package main

import (
    "errors"
    "fmt"
    "os"
)

func main() {
    if err := doError(); err != nil {
        fmt.Println("err", err.Error())
        os.Exit(1)
    }
    fmt.Println("(o・∇・o)終わりだよ～") // ここにはこない
}
func doError() error {
    return errors.New("(*>△<)<ナーンナーンっっ")
}
