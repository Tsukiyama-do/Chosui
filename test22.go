package main

import (
    "fmt"
    "sort"
)

func main(){
   i := []int{ 5, 230,33,122213,323233, 23232,6344323, 3243223}

   sort.Sort(sort.IntSlice(i))
   fmt.Println(i)

   j := []int{ 3233, 230,33,127213,323233, 93232,6344323, 3243223}

   sort.Sort(sort.Reverse(sort.IntSlice(j)))
   fmt.Println(j)

}
