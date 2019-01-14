package main

import (
    "errors"
    "fmt"
    "os"
)

func main() {
    if err := doError(); err != nil {
        err = fmt.Errorf("MOBIUS Err :  %s", err)
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println("(o・∇・o)終わりだよ～") // ここにはこない
}
func doError() error {
    return errors.New("Hirose (*>△<)<ナーンナーンっっ")
}
