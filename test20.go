package main

import (
    "golang.org/x/text/width"
    "fmt"
)

func showUnicode(s string){
  w := width.Widen.String(s)
  fmt.Printf("%U: %s\n", []rune(s), s)
  fmt.Printf("%U: %s\n", []rune(w), w)
}

func main() {
    s_simple := "ﾊﾋﾌﾍﾎ"
    s_dakuten := "ﾊﾞﾋﾞﾌﾞﾍﾞﾎﾞ"
    s_handakuten := "ﾊﾟﾋﾟﾌﾟﾍﾟﾎﾟ"
    s_longvow := "ﾊﾛｰ"
    showUnicode(s_simple)
    showUnicode(s_dakuten)
    showUnicode(s_handakuten)
    showUnicode(s_longvow)

}
