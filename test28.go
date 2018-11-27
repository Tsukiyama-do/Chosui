package main
import (
  "fmt"
  "sort"
)

type IntSlice []int


func (p IntSlice) Less(i, j int) bool { return p[i] > p[j] }

func main() {

  nums := []int{4, 3, 2, 10, 8}

  sort.Ints(IntSlice(nums))
  fmt.Println(nums)



}
