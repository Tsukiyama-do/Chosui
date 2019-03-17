package main

import (

  "fmt"
  "sort"
)

type Change struct{
  user string
  language string
  lines int
}

type lessFunc func(f1, f2 *Change) bool

type multiSorter struct {

  changes []Change
  less []lessFunc
}

func (ms *multiSorter) Sort(changes []Change ) {

  ms.changes = changes
  sort.Sort(ms)
}

func OrderedBy(less ...lessFunc) *multiSorter {

  return &multiSorter{
      less: less,
  }
}

func (ms *multiSorter) Len() int {
  return len(ms.changes )

}
func (ms *multiSorter) Swap(i,j int)  {
  ms.changes[i], ms.changes[j] = ms.changes[j], ms.changes[i]
}

func (ms *multiSorter) Less(i, j int ) bool {
  p, q := &ms.changes[i], &ms.changes[j]
    var k int
    for k = 0; k < len(ms.less)-1; k++ {
      less := ms.less[k]
      switch {
      case less(p, q):
              return true
      case less(q, p):
              return false
      default :
          //  p == q
      }
    }
    return ms.less[k](p,q)
}

var changes = []Change{
  {"gri", "Go", 120},
  {"ken", "C", 130},
  {"glenda", "Go", 150},
  {"rsc", "Go", 250},
  {"r", "Go", 100},
  {"r", "C", 140},
  {"gri", "Smalltalk", 80},
  {"ken", "Go", 150},


}

func main() {

  user := func(j1, j2 *Change) bool {
      return j1.user < j2.user
  }

  language := func(j1, j2 *Change) bool {
      return j1.language < j2.language
  }
  increasingLines := func(j1, j2 *Change) bool {
      return j1.lines < j2.lines

  }
  decreasingLines := func(j1, j2 *Change) bool {
      return j1.lines > j2.lines

  }
    OrderedBy(user).Sort(changes)
    fmt.Println("By user : \n")
    for idx, d := range changes { fmt.Printf("No. %d, %s, %s, %d \n", idx, d.user, d.language, d.lines) }
    OrderedBy(user, increasingLines).Sort(changes)
    fmt.Println("By user, <lines  \n")
    for idx, d := range changes { fmt.Printf("No. %d, %s, %s, %d \n", idx, d.user, d.language, d.lines) }
    OrderedBy(user, decreasingLines).Sort(changes)
    fmt.Println("By user, >lines  \n")
    for idx, d := range changes { fmt.Printf("No. %d, %s, %s, %d \n", idx, d.user, d.language, d.lines) }
    OrderedBy(language, increasingLines).Sort(changes)
    fmt.Println("By language, <lines  \n")
    for idx, d := range changes { fmt.Printf("No. %d, %s, %s, %d \n", idx, d.user, d.language, d.lines) }
    OrderedBy(language, increasingLines, user).Sort(changes)
    fmt.Println("By language, <lines, users  \n")
    for idx, d := range changes { fmt.Printf("No. %d, %s, %s, %d \n", idx, d.user, d.language, d.lines) }


}
