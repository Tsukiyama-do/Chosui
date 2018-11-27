package main

import (
  "fmt"
  "sort"
)

type Food struct {
  Name string
  Price int
}

type Foods []Food

func (t Foods) Len() int {
  return len(t)

}

func (t Foods) Less(i, j int) bool {
  return t[i].Price >= t[j].Price
}

func (t Foods) Swap(i, j int) {
  t[i], t[j] = t[j], t[i]
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

  sort.Sort(Foods(foods))

  for _, food := range foods {
    fmt.Printf("商品名：　%s, 価格（税込）: %d\n", food.Name, food.Price)
  }

}
