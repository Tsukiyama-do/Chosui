package main

import (

   "log"
//   "hash"
//   "github.com/Logibox/rollinghash"
  "os"
)


func main(){

  _, err := os.Open("/tmp/hogehoge.txt")
  if err != nil {
      // エラー時の処理
      log.Printf("Error is " + err.Error())
      log.Fatal(err)
    }

}
