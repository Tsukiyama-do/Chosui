package main

import "fmt"
// import "reflect"

type Foo interface {
    Echo() string
}
type Bar interface {
    Unko() string
}
type Fooer struct {
}
func (f Fooer)Echo() string {
    return "implements Echo"
}
func (f Fooer)Unko() string {
    return "implements Unko"
}



func main() {
    f := Fooer{}

    fooer, ok := interface{}(f).(Foo)
    fmt.Printf("%T\n", fooer)
    fmt.Printf("%T\n", ok)
    fmt.Println(fooer.Echo(), ok)

    barer, ng := interface{}(f).(Bar)
    fmt.Printf("%T\n", barer)
    fmt.Printf("%T\n", ng)
    fmt.Println(barer.Unko(), ng)
}
