package main
import (
  "fmt"
)

type as_i interface{
  whereis() string
}

type as_j interface{
}

type as_k interface{
  myname() string
}

type hr_t struct{
  address string
  tel string
}

type hr_s string

type hr_i int

func (v_hr *hr_t) whereis() string{
  return fmt.Sprintf("You are : %s %s ", v_hr.address, v_hr.tel)
}

func (v_s hr_s) myname() string{
  return fmt.Sprintf("My name is %s.", v_s)

}

func checkType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("私はintです")
	case float64:
		fmt.Println("私はfloat64です")
	case string:
		fmt.Println("私はstringです")
  case hr_s:
		fmt.Println("私はhr_sです")
  case hr_t:
		fmt.Println("私はhr_tです")
  case *hr_t:
		fmt.Println("私は*hr_tです")
	default:
		fmt.Println("私はそれ以外です")
	}
}


func main() {

//    st_v := hr_s("Koyasu")
    place_info := hr_t{address : "Nisikasai Nichome", tel : "070-8888-9923"}

    var inp as_i
//    inp = st_v
    inp = &place_info

    fmt.Printf("Hi, this is result : %s \n",inp.whereis())
    checkType(inp)

//    inp.myname()

    var inq as_j
    hename := hr_s("00005")
    inq = hename
    fmt.Printf("Hoi, this is result : %s \n", inq)
    checkType(inq)

    var inr as_k
//    yname := hr_s("0000006")
//    inr = yname
    inr = hr_s("0000006")     //   この型で、この値と決めて、インターフェイスに値を入れる

    fmt.Printf("Hum, this is result : %s \n", inr.myname())    // インターフェイスから、メソッドを呼べる
    checkType(inr)

}
