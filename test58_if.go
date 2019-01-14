package main

import "fmt"

func main() {
	var i interface{}

	i = 3
	checkType(i)
	i = 3.5
	checkType(i)
	i = "aaaaaa"
	checkType(i)

	type MyType uint
	i = MyType(4)
	checkType(i)

}

func checkType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("私はintです")
	case float64:
		fmt.Println("私はfloat64です")
	case string:
		fmt.Println("私はstringです")
	default:
		fmt.Println("私はそれ以外です")
	}
}
