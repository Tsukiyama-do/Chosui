package main

import (
  "fmt"

)

func main() {

//  var ib int
//  ib = 1

  var listint = []int{1,1}
    for i := 2; i < 20;i++ {
      listint = append(listint, listint[i - 1] + listint[i - 2])
    }
    printSlice(listint)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
