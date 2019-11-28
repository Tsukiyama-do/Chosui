package main

import (
 "fmt"
 "strconv"
)

type Stringer interface {
 String() string
}

type Binary uint64

func (i Binary) String() string {
 return strconv.Itob64(i.Get())
}

func (i Binary) Get() uint64 {
 return uint64(i)
}

func main() {
 b := Binary(200)
 s := Stringer(b)
 fmt.Println(s.String())
}
