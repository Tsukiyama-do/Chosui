package main

import (
    "log"
)

func init() {
    log.SetFlags(log.Lshortfile)
}

func add(a []int) {
    a = append(a, 4)           // 配列自体は更新されないはず
    log.Printf("%p, %v", a, a) //  0x18244010, [1 2 3 4]
}

func main() {
    a := make([]int, 3, 5)
    a[0], a[1], a[2] = 1, 2, 3 // まだ cap に余裕がある
    log.Printf("%p, %v", a, a) //  0x18244010, [1 2 3]
    add(a)
    a = a[:cap(a)]
    log.Printf("%p, %v", a, a) // 0x18244010, [1 2 3]
}
