package main
import (
  "fmt"
)
type Person struct{
  first_name string
  last_name string
  dateofbirth string
  sex string
  country string
}

type Employee struct{
  Person
  section string
  work_years int

}

func (p *Person) String() string {
  return fmt.Sprintf("Name : %s %s, M/F : %s", p.first_name, p.last_name, p.sex)
}

func (p *Person) CountryCode() string {
  var code string
  if p.country == "JAPAN" {
    code = "JP"
  }
  if p.country == "USA" {
    code = "US"
  }
  if p.country == "CHINA" {
    code = "CH"
  }

  return code

}

func (e *Employee) Sections() string {
  var code string
  if e.section == "FSDD1" {
    code = "Financial System Development Division 1"
  }
  if e.section == "FSDD2" {
    code = "Financial System Development Division 2"
  }
  if e.section == "FSDD3" {
    code = "Financial System Development Division 3"
  }

  return code


}

type Men interface {
  CountryCode() string
  Sections() string
}


func main() {

  narita := &Employee{Person{"Mamoru", "Narita", "20100909","Male","JAPAN"}, "FSDD1", 5}
  nagata := &Employee{Person{"Sho", "Nagata", "20100109","Female","USA"}, "FSDD2", 20}
  mizuno := &Employee{Person{"Sanae", "Mizuno", "20000209","Female","JAPAN"}, "FSDD1", 10}

  var i Men
//  i = narita
  fmt.Printf("Narita data is : %s\n", narita)
  fmt.Printf("Narita Country is : %s\n", narita.CountryCode())
  fmt.Printf("Narita Section is : %s\n", narita.Sections())

  i = nagata
  fmt.Printf("Nagata data is : %s\n", i.CountryCode())
  fmt.Printf("Nagata Section is : %s\n", i.Sections())
  i = mizuno
  fmt.Printf("Mizuno data is : %s\n", i.CountryCode())
  fmt.Printf("Mizuno Section is : %s\n", i.Sections())

}
