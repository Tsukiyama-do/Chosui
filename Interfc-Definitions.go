


https://www.blog.umentu.work/%E3%80%90%E5%AD%A6%E7%BF%92golang%E3%82%B7%E3%83%AA%E3%83%BC%E3%82%BA%E3%80%91%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%BC%E3%83%95%E3%82%A7%E3%83%BC%E3%82%B9%E3%81%A7%E3%83%9D%E3%83%AA%E3%83%A2%E3%83%BC/

インターフェースを使う手順は次の通り。

構造体を定義する
構造体のメソッドを定義する
インターフェースに構造体のメソッド名を入れる
main関数内で構造体型の変数を定義する。
main関数内でインターフェース型の変数を定義し、構造体型の変数を代入する
インターフェース型の変数から、構造体のメソッドを呼び出す



package main

import (
    "fmt"
)

type Human interface {
    // インターフェースにはTalkメソッドとWalkメソッドのみいれてある
    Talk()
    Walk()
    // Run()
}

// 日本人構造体
type Japanese struct {
    name string
    age int
}

// 米国人構造体
type American struct {
    name string
    age int
}

// 日本人用Talkメソッド
func (j Japanese) Talk(){
    fmt.Printf("%sです。%dです。
", j.name, j.age)
}

// 日本人用Walkメソッド
func (j Japanese) Walk(){
    fmt.Println("トコトコ")
}

// 日本人用Runメソッド
func (j Japanese) Run(){
    fmt.Println("シュタタタ")
}

// 米国人用Talkメソッド
func (a American) Talk(){
    fmt.Printf("I'm %s.I'm %d years old.
", a.name, a.age)
}

// 米国人用Walkメソッド
func (a American) Walk(){
    fmt.Println("tokotoko")
}

// 米国人用Runメソッド
func (a American) Run(){
    fmt.Println("shutatata")
}

func main(){

    j := Japanese{"umentu", 5}
    // 構造体の値をインターフェースに設定
    var h Human = j
    h.Talk()
    h.Walk()
    // h.Run() はHumanに入っていないから実行できない

    a := American{"John", 100}
    // 構造体の値をインターフェースに設定
    var h2 Human = a
    h2.Talk()
    h2.Walk()
    // h2.Run() はHumanに入っていないから実行できない
}
