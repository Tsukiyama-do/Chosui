package main

import (
  "fmt"
  "strings"
)

/*
var obj1, obj2, obj3, obj4, obj5 interface{}

obj1 = 123                                                              // int
obj2 = "str"                                                            // string
obj3 = []string{"linux", "windows", "android"}                          // slice
obj4 = make(chan string)                                                // channel
obj5 = func(val int) string { return fmt.Sprintf("number is %d", val) } // function

*/

func printIf(src interface{}) {
    if value, ok := src.(int); ok {
        fmt.Printf("parameter is integer. [value: %d]\n", value)
        return
    }

    if value, ok := src.(string); ok {
        value = strings.ToUpper(value) // 対象がstring型なのでstringを引数に取る関数が実行できる
        fmt.Printf("parameter is string. [value: %s]\n", value)
        return
    }

    if value, ok := src.([]string); ok {
        value = append(value, "unknown") // 対象がsliceなのでAppendができる
        fmt.Printf("parameter is slice string. [value: %s]\n", value)
        return
    }

    fmt.Printf("parameter is unknown type. [valueType: %T]\n", src)
}

func printIIf(src interface{}) {
    switch value := src.(type) {
    case int:
        fmt.Printf("parameter is integer. [value: %d]\n", value)
    case string:
        value = strings.ToUpper(value) // 対象がstring型なのでstringを引数に取る関数が実行できる
        fmt.Printf("parameter is string. [value: %s]\n", value)
    case []string:
        value = append(value, "<不明>") // 対象がsliceなのでAppendができる
        fmt.Printf("parameter is slice string. [value: %s]\n", value)
    default:
        fmt.Printf("parameter is unknown type. [valueType: %T]\n", src)
    }
}

func main() {
    printIf(12)
    printIf("hello")
    printIf([]string{"cat", "dog"})
    printIf([2]string{"hello", "world"})

    printIIf(12)
    printIIf("hello")
    printIIf([]string{"cat", "dog"})
    printIIf([2]string{"hello", "world"})
}
