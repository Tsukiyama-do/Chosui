package main

import (
  "fmt"
)

type Duck interface {
    Quack() string
}

type duck struct {}

func NewDuck() Duck {
    return &duck{}
}

func (d *duck) Quack() string {
    return "QUUAAAACCCCKK!!!"
}

func main() {

  duck := NewDuck()
   if x := duck.Quack(); len(x) < 1 {
       fmt.Println("duck doesn't quack")
   }else {
      fmt.Println(x)

  }

}
