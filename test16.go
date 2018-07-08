package main

import (
  "fmt"

)

func main() {
  ss := [4]int{ 5, 6, 2, 200}
  tt := []int{55, 33, 49,33, 10, 5}

  st := make([]int, 10, 20)

  st = ss[1:3]
  fmt.Println(st)
  fmt.Println(tt)
//   st = append(st, tt...)
  st = append(st,tt[2])
//  st[3] = tt[3]

  fmt.Println(st)

 var jz1 *int
 var jz2 int
 var jz9 *string

 jline := [5]byte{'a','b','c','d','\n'}
 jline2 := &jline


 jz2 = 3
 jz1 = &jz2

 jz8 := "Hirose Hospital"
 jzA := "Hirose Supermarket"

 jz9 = &jz8

 fmt.Printf("Int is %d\n", *jz1)
 fmt.Printf("Int is %d\n", jz2)

 fmt.Printf("Int is %s\n", *jz9)
 fmt.Printf("Int is %s\n", jz8)

 fmt.Println(jline)
 fmt.Println(jline2)

 jz2 = 99
 jz9 = &jzA
 fmt.Printf("Int is %d\n", *jz1)
 fmt.Printf("Int is %d\n", jz2)
 fmt.Printf("Int is %s\n", *jz9)
 fmt.Printf("Int is %s\n", jz8)



}
