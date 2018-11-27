package main

import (
    "fmt"
    "sort"
)

type Item struct {
    id    string
    score int
}

func add(items []Item, item Item) []Item {
    fmt.Println(items, "<=", item)

    i := sort.Search(len(items), func(i int) bool {
        return items[i].score < item.score ||
            (items[i].score == item.score && items[i].id > item.id)
    })

    items = append(items, Item{})
    copy(items[i+1:], items[i:])
    items[i] = item

    return items
}

func main() {
    items := []Item{}
    items = add(items, Item{"A", 100})
    items = add(items, Item{"C", 200})
    items = add(items, Item{"B", 200})
    items = add(items, Item{"D", 50})
    items = add(items, Item{"E", 100})
    items = add(items, Item{"F", 40})
    items = add(items, Item{"G", 200})
    items = add(items, Item{"H", 220})
    fmt.Println(items)
}
