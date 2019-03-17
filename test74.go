package main
import (

  "fmt"

  "golang.org/x/text/language"
  "golang.org/x/text/language/display"
)

var uPrefs = []language.Tag{
  language.Make("gsw"),  // Swiss German
  language.Make("fr"),
}


var servLangs = []language.Tag{
  language.SimplifiedChinese,    // simple chinese
  language.TraditionalChinese,    // simple chinese
  language.English,
  language.French,
  language.Japanese,

}

var matcher = language.NewMatcher(servLangs)

func main() {

  var uPref2 = []language.Tag{}

  sTag, err := language.Parse("my")
  if err != nil {
      fmt.Println("fr is invalid.")
  } else {
    fmt.Printf("v : %v \n", sTag)
      uPref2 = append(uPref2, sTag)
  }


  tag, index, confidence := matcher.Match(uPref2...)
  fmt.Printf("best match : %s (%s) index=%d confidence=%v \n", display.English.Tags().Name(tag), display.Self.Name(tag), index,confidence)  // best match : Germanl

}
