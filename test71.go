package main

import (

  "fmt"
)

type Dog interface {
  String() string
}

type dog struct{}

func CreateDog() Dog {
  return &dog{}

}

func (d *dog) String() string {

  return "Wanwan!"
}

func main() {

  Lov := CreateDog()
  fmt.Println("Love is ", Lov.String())


}
