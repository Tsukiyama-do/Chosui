package main

import (
    "fmt"
//    "strconv"
)
//   https://dev.classmethod.jp/go/golang-6/

//空インターフェイス
type Element interface{}


func main() {
  //適当な値をセット
  var element Element = 1000

  //どの型にマッチするかチェック
  if value, ok := element.(int); ok {
    fmt.Printf("int value: %d \n", value)
  } else if value, ok := element.(string); ok {
    fmt.Printf("string value: %s \n", value)
  } else {
    fmt.Println("other \n")
  }

}
