package main

import (
  "fmt"
  "sort"
)

type Food struct {
  Name string
  Price int
}

func main() {

  var foods = []Food{
    {"鰻丼並", 750},
    {"牛丼並", 380},
    {"牛丼頭", 450},
    {"鰻丼大", 890},
    {"牛丼大", 550},
    {"牛丼特", 550},

  }


  for _, food := range foods {
    fmt.Printf("%+v\n", food)
  }

  sort.Slice(foods, func(i, j int) bool{
    return foods[i].Price < foods[j].Price
  })

  for _, food := range foods {
    fmt.Printf("商品名：　%s, 価格（税込）: %d\n", food.Name, food.Price)
  }

}
