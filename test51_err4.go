package main

import (
    "fmt"
    "os"
)

// エラー処理用の構造体
type HRError struct {
    Msg string
    Code int
}
// MyError構造体にerrorインタフェースのError関数を実装
func (err *HRError) Error() string {
    return fmt.Sprintf("HR Error : %d : %s ",  err.Code, err.Msg)
}

func main() {
    if err := doError(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println("(o・∇・o)終わりだよ～") // ここにはこない
}
func doError() error {
    return &HRError{Msg:"Oh my goodness!", Code:20202}
}
