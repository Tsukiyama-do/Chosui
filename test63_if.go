package main
import (
  "fmt"
)

type Coffee struct{
  name string
  country string
  unitprice int
}

func (c *Coffee) String() string{
    return fmt.Sprintf("Name : %s, Country : %s, Price : JPY %d ",c.name, c.country, c.unitprice)
}

func (c *Coffee) VolumePrice() string{
    return fmt.Sprintf("Name : %s, Country : %s, Price : JPY %d ",c.name, c.country, c.unitprice*100)
}


func main() {
  cb := &Coffee{name : "Colombia", country : "Africa", unitprice : 20 }

  fmt.Println(cb)
  fmt.Println(cb.VolumePrice())

}
