package main

import (
    "fmt"
    "sort"
)




func main() {
  line  := []float64 { 2393202123, 832923013203,9329322303232, 3288230923,32322 }
  fmt.Println(line)
  sort.Sort(sort.Float64Slice(line))
  fmt.Println(line)
  sort.Sort(sort.Reverse(sort.Float64Slice(line)))

  fmt.Println(line)

}
