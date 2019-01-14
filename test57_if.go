package main

import "fmt"

func main() {
 var i interface{}
 i = 4
 fmt.Println(i)             //4
 i = 4.5
 fmt.Println(i)             //4.5
 i = "文字列だってはいるんだ"
 fmt.Println(i)             //文字列だってはいるんだ
}
