package main



import "fmt"

var table [][]int

func loadTable(a *[]int) {
    table = make([][]int, len(*a))
    for _, e := range *a {
        h := hash(e, len(*a))
        fmt.Printf(" value is  %d and the hash is %d\n", e, h)
        if table[h] == nil {
            table[h] = make([]int, 0)
        }
        table[h] = append(table[h], e)
    }

    for i := 0 ; i < 10 ; i++ {
      fmt.Printf("table value is %+v\n", table[i])
    }
}

func hash(key int, size int) int {
    var h int
    for i := 0; i < key+1; i++ {
        h = (h*137 + i) % size
    }
    return h
}

func search(a *[]int, t int) bool {
    h := hash(t, len(*a))
    fmt.Printf("hash is %d\n", h)
    list := table[h]
    if list == nil {
        return false
    }
    for _, v := range list {
        if v == t {
            return true
        }
    }
    return false
}

func main() {
    a := []int{4, 19, 3, 8, 5, 16, 7, 2, 90, 10}
    loadTable(&a)
    fmt.Println(a)
    fmt.Printf("0:%t\n", search(&a, 0))
    fmt.Println(a)
    fmt.Printf("2:%t\n", search(&a, 2))
    fmt.Println(a)
    fmt.Printf("12:%t\n", search(&a, 12))
}
