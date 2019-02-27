package main

import (

  "fmt"
)

type Whale interface {
  String() string
  Jump() int
  Swim()
}

type whale struct{
  name string
  long int
  weight int
  color string
  food string
  aveage int
}

func (a *whale)String() string {

  return fmt.Sprintf("Whale : %s, color: %s", a.name, a.color)

}

func (a *whale) Jump() int {

  return a.long * a.weight

}

func (a *whale) Swim() {
  fmt.Printf("Swim faster by %d \n", a.weight * 2 )

}

func NewWhale(a string, b int, c int, d string, e string, f int) Whale {
  return &whale{a, b, c, d, e, f }

}

func main() {

//  var Jimbe Whale
//  Jimbe = &whale{ "Jimbe zame" , 5, 3, "blue", "Iwashi", 20 }

  Jimbe := NewWhale("Jim", 6, 8, "blue", "Aji", 30)

  fmt.Println(Jimbe.String())
  fmt.Printf("Jump is %d \n", Jimbe.Jump())
  Jimbe.Swim()
}
