package main

import (
  "fmt"

)


type Macmenu struct{
  Name string
  Price int
}






func main() {


var items = []Macmenu {
  {"Ice Coffee small", 100},
  {"Apple pie", 100},
  {"Burgar", 100},
  {"Cheese burgar", 120},
}

var p *[]Macmenu
   p = &items

   for _, e := range *p {
     fmt.Printf("Item name is %s, and the price is %d\n",e.Name, e.Price )
   }

  fmt.Printf("The pointer is %p\n",p )
  fmt.Printf("%+v\n",*p)


}
