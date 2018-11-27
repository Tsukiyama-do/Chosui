package main

import (
    "fmt"
    "sort"
)

type Student struct {
  LastName string
  FirstName string
  Sex string
  Department int
  Position int
  Age int

}


func (s Student) String() string {
  return fmt.Sprintf("%s:%s:%s:%d:%d:%d", s.LastName, s.FirstName,s.Sex,s.Department,s.Position,s.Age)
}

type ByLastName []Student

func (x ByLastName) Len() int {
  return len(x)
}

func (x ByLastName) Swap(i, j int) {
  x[i], x[j] = x[j], x[i]
}

func (x ByLastName) Less(i, j int) bool {
  return x[i].LastName < x[j].LastName
}


type ByFirstName []Student

func (x ByFirstName) Len() int {
  return len(x)
}

func (x ByFirstName) Swap(i, j int) {
  x[i], x[j] = x[j], x[i]
}

func (x ByFirstName) Less(i, j int) bool {
  return x[i].FirstName < x[j].FirstName
}


type ByAge []Student

func (x ByAge) Len() int {
  return len(x)
}

func (x ByAge) Swap(i, j int) {
  x[i], x[j] = x[j], x[i]
}

func (x ByAge) Less(i, j int) bool {
  return x[i].Age < x[j].Age
}

func main() {
  people := []Student {
    {"Hirose", "Jun","Male",2,2,22},
    {"Akiyoshi", "Sanae","Female",2,2,32},
    {"Shimazu", "Yoichi","Male",2,1,15},
    {"Kawakami", "Sanae","Female",2,1,32},
    {"Yoshida", "Ichiro","Male",2,1,45},
    {"Wada", "Jinmo","Male",2,1,30},
  }
  for _, item := range people {
    fmt.Printf("Last name: %s, First name : %s, M/F : %s\n", item.LastName, item.FirstName, item.Sex)
  }
  sort.Sort(ByLastName(people))
  for _, item := range people {
    fmt.Printf("Last name: %s, First name : %s, M/F : %s\n", item.LastName, item.FirstName, item.Sex)
  }

  sort.Sort(ByFirstName(people))
  for _, item := range people {
    fmt.Printf("Last name: %s, First name : %s, M/F : %s\n", item.LastName, item.FirstName, item.Sex)
  }
  sort.Sort(ByAge(people))

  for _, item := range people {
    fmt.Printf("Last name: %s, First name : %s, M/F : %s\n", item.LastName, item.FirstName, item.Sex)
  }


}
