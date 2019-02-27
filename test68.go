package main

import (
  "fmt"
  "sort"
)

type Fsdd2 string

var s_name = []Fsdd2{
     "Hirose" ,
     "Maeda" ,
     "Nakano" ,
     "Akaogi" ,
     "Tmomi" ,
}

type SrtFsdd2 struct{
  name []Fsdd2
}

func (a *SrtFsdd2) Len() int {
    return len(a.name)
}

func (a *SrtFsdd2) Swap(i, j int ) {
    a.name[i], a.name[j] = a.name[j], a.name[i]
}

func (a *SrtFsdd2) Less(i, j int ) bool {
    return a.name[i] <= a.name[j]
}


func main() {

  nameg := &SrtFsdd2{
    name : s_name,
  }

  sort.Sort(nameg)
  fmt.Println("Fsdd2 is ", nameg)
  for idx, j := range nameg.name {

    fmt.Printf("Item %d, %s \n",idx, j)

  }

}
