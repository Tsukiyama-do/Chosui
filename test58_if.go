package main

import "fmt"

type MyType uint


func main() {
	var i interface{}

	i = 3
	checkType(i)
	i = 3.5
	checkType(i)
	i = "aaaaaa"
	checkType(i)

	i = MyType(4)
	checkType(i)

}

func checkType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Printf("私はintです %d\n", i)
	case float64:
		fmt.Printf("私はfloat64です %f\n", i)
	case string:
		fmt.Printf("私はstringです %s\n", i)
	case MyType:
		fmt.Printf("私はMyTypeです %d\n",i)

	default:
		fmt.Printf("私はそれ以外です")
	}

}
