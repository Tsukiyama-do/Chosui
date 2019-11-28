package main

import (
    "fmt"
    "time"
)

func main() {
    q := make(chan string, 5)

    go func() {
        time.Sleep(3 * time.Second)
        q <- "foo"
    }()
    go func() {
        time.Sleep(5 * time.Second)
        q <- "bar"
    }()

    // ほぼ同時に2つトリガが発生するのが分かっている
    var cmds []string
    cmds = append(cmds, <-q) // まずは一つ受信する
wait_some:
    for {
        select {
        case <-time.After(1 * time.Second):
            // 1秒過ぎたらもう受け付けないよ
            break wait_some
        case cmd := <-q:
            // 1秒以内なら一緒に処理しちゃうよ
            cmds = append(cmds, cmd)
        }
        fmt.Println("Hi")
    }
    for _, cmd := range cmds {
        fmt.Println(cmd)
    }
}
