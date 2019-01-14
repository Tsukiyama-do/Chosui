package main

import (
  "fmt"

)

type MyInterface interface{
  adder(x int)(int)
}
type Struct1 struct{
  aa int
}
type Struct2 struct{
  bb int
}
func(s1 Struct1)adder(x int)(int){
   return s1.aa + x
}

func main() {
  ss1 :=  Struct1{aa : 2}
//  ss2 :=  Struct2{bb : 2}
  var ia1 MyInterface
//  var ia2 MyInterface
  ia1 = ss1
//  ia2 = ss2

  fmt.Printf(" Values are %d\n", ia1.adder(2))
//  fmt.Printf(" Values are %d\n", ia2.adder(2))
}
