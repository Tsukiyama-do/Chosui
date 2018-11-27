package main

import (
    "fmt"
    "math/rand"
)

func quickSort(a *[]int, left int, right int) {
    var  pivot_l,pivot_r int

    if right > left {
        pivot_l, pivot_r = partition(a, left, right)
    }
    if left < (pivot_l) {
        quickSort(a, left, pivot_l)
    }
    if ( pivot_l+1 ) < right {
        quickSort(a, pivot_l+1, right)
    }
    if pivot_r > 0 {
      pivot_r = pivot_r +1
    }
}

func partition(a *[]int, left int, right int) (int, int) {
//    tmp_left := left
//    var wk int
    pivotItem := (*a)[(left + right) / 2 ]
    fmt.Println(pivotItem)
    fmt.Println(*a)

    for ; left < right ; {
        for (*a)[left] <= pivotItem {
            left++
            if left >= right { break }
        }
        if left >= right { break }
        for (*a)[right] > pivotItem {
            right--
            if left >= right { break }
        }
        if left >= right { break }

        if left < right {
            (*a)[right], (*a)[left] = (*a)[left], (*a)[right]
        }
    }
    fmt.Printf("Left, Right are  %d %d\n",left, right)
    return left, right
}

func main() {
    var ary []int
    size := 50
    for i := 0; i < size; i++ {
        ary = append(ary, rand.Intn(size-1))
    }
    fmt.Println(ary)
    quickSort(&ary, 0, len(ary)-1)
    fmt.Println(ary)
}
