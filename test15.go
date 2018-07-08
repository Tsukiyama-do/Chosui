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

type Product struct {
    id    uint
    name  string
    price uint
    PR    PRStatement
}

type PRStatement func() string

// PRStatementという関数をレシーバにして対象インターフェースのメソッドを定義
func (pr PRStatement) description() string {
    return pr()
}

func main() {
    p1 := &Product{id: 1, name: "Golang PC", price: 10000}
    p1.PR = func() string {
        return fmt.Sprintf("この %s は、値段が%d円なのでとてもお買い得です", p1.name, p1.price)
    }
    printDescription(p1.PR)

    p2 := &Product{id: 2, name: "リンゴ", price: 100}
    p2.PR = func() string {
        return fmt.Sprintf("この %s は、とても美味しいです", p2.name)
    }
    printDescription(p2.PR)
}
