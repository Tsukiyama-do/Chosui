package main

import (
    "time"
)

func main() {
    q := make(chan struct{})

    go func() {
        // 重たい処理
        time.Sleep(3 * time.Second)
        q <- struct{}{}
    }()

    // q に何か入るまで待つ
    <-q
}
