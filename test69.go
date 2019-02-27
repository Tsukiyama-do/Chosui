package main

import (

    "fmt"
)

type MyInterface interface{
  adder(x int)(int)
}
type Struct1 struct{
  num int
}
type Struct2 struct{
  num int
}
func(s1 Struct1)adder(x int)(int) {
   return  s1.num + x
}

func main() {

  var i MyInterface
  i = Struct1{ 5 }
//  j = Struct2{ 10 }

  fmt.Printf("Result : %d",i.adder(8) )
  fmt.Printf("Result2 : %d",i.num )
}
