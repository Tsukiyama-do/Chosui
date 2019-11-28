package main

import (
  "fmt"
)

type sakana interface{
  shop_name() string
  aji() string
  saba() string
  ika() string

}

type yasai interface{
  shop_name() string
  tomato() string
  ninjin() string
}

type niku interface{
  shop_name() string
  pig() string
  beef() string

}

type kutsu interface{
  shop_name() string
  sandal() string
  boot() string
}


type shop_info struct{
  Name string


}

func (vt *shop_info) shop_name() string {
    return vt.Name
}

func (vt *shop_info) aji() string {
    return "Aji Aji"
}

func (vt *shop_info) saba() string {
    return "Saba Saba"
}

func (vt *shop_info) ika() string {
    return "Ika"
}

func checkType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("私はintです")
	case float64:
		fmt.Println("私はfloat64です")
	case string:
		fmt.Println("私はstringです")
  case *shop_info:
		fmt.Println("私は*shop_infoです")
	default:
		fmt.Println("私はそれ以外です")
	}
}

func fish_check(a interface{}) {
  switch a.(type) {
	case int:
		fmt.Println("私はintです")
	case float64:
		fmt.Println("私はfloat64です")
	case *shop_info:
		fmt.Println("私はshop_infoです")
    fmt.Println(a.(*shop_info).Name)    //  これぞインターフェイスのアクセス方法です！
    if aj, ok := a.(*shop_info); ok {    //  これがインターフェイスからのキャストらしい
      fmt.Println("casted! ", aj.Name)
      }     
  case yasai:
		fmt.Println("私はyasaiです")
	default:
		fmt.Println("私はそれ以外です")
	}

}

func main () {

  vs := shop_info{Name: "Higashi kanagawa"}
  var s_sakana sakana = &vs

  fmt.Printf("Hi, this is result : %s \n",s_sakana.shop_name())
  checkType(s_sakana)
  fmt.Printf("Hi, this is result2 : %s \n",s_sakana.aji())

  fish_check(s_sakana)

}
