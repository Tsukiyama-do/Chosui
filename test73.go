package main

import (

    "fmt"
)

type MyInterface interface{
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

  var i, j MyInterface
  i = Struct1{ 5 }
  j = Struct2{ 10 }

  fmt.Printf("Result : %v \n",i )

	g, ok := i.(Struct1) // v が Get() を実装しているか調べる
	if ok {
    fmt.Printf("Result if : %d \n", g.num )
	} else {
    fmt.Printf("Error")
	}

  fmt.Printf("Result2 : %v \n",j )
//  fmt.Printf("Result2: %d \n",j.(map[string]interface{}["num"].(int)))

}
