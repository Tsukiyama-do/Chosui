package main

import (
  "fmt"

)

type describe interface {
    description() string
}

func printDescription(d describe) {
    fmt.Printf("Description: %s\n", d.description())
}

type MyInt int

// 組み込み型の別名の型をレシーバにして対象インターフェースのメソッドを定義
func (i MyInt) description() string {
    return fmt.Sprintf("MyInt is actually int. value is %d", i)
}

func main() {
    var val1 MyInt = 500
    printDescription(val1)
}
