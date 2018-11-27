package main

import (
//  "fmt"
//  "sort"
  "log"
  "time"
)

func init() {
    log.SetFlags(log.Lshortfile)
}

type Food struct {
  Name string
  Kana string
  Price int
}

type Fline struct {
  Name string
  Kana string
  Price int
  Nameb []byte
  Kanab []byte
}

func insertsort(a []Fline) []Fline {


  log.Println(time.Now())           // => "2015-05-05 07:23:30.757800829 +0900 JST"


  for i := 1 ; i < len(a)  ; i++  {     //  insert sort logic
    for j := i - 1 ; j >= 0 ; j-- {    //  insert sort logic
        if a[j].Price > a[j+1].Price {
          a[j], a[j+1] = a[j+1], a[j]
        }
    }
  }

  log.Println(time.Now())           // => "2015-05-05 07:23:30.757800829 +0900 JST"
  return a

}

func main() {

  var foods = []Food{
    {"鰻丼並","ｳﾅﾄﾞﾝﾅﾐ", 750},
    {"牛丼並","ｷﾞｭｳﾄﾞﾝﾅﾐ", 380},
    {"牛丼頭","ｷﾞｭｳﾄﾞﾝｱﾀﾏ", 450},
    {"鰻丼大","ｳﾅﾄﾞﾝﾀﾞｲ", 890},
    {"牛丼大","ｷﾞｭｳﾄﾞﾝﾀﾞｲ", 550},
    {"ペペロンチーノ","ﾍﾟﾍﾟﾛﾝﾁｰﾉ", 800},
    {"牛鍋膳","ｷﾞｭｳﾅﾍﾞｾﾞﾝ", 700},
    {"ビックマックセット","ﾋﾞｯｸﾏｯｸｾｯﾄ",    452},
    {"ソーセージマフィン","ｿｰｾｰｼﾞﾏﾌｨﾝ",  200},
    {"エッグマフィン","ｴｯｸﾞﾏﾌｨﾝ",  300},
    {"コーヒー","ｺｰﾋｰ",  150},
    {"ポテトM","ﾎﾟﾃﾄｴﾑ",  200},
    {"ピスタチオ","ﾋﾟｽﾀﾁｵ",  1220},
    {"麻婆丼","ﾏｰﾎﾞｰﾄﾞﾝ",  900},
    {"刺し身定食","ｻｼﾐﾃｲｼｮｸ",  1200},
    {"納豆定食","ﾅｯﾄｳﾃｲｼｮｸ",  290},
    {"とんかつ定食","ﾄﾝｶﾂﾃｲｼｮｸ",  1100},
    {"つけ麺ラーメン","ﾂｹﾒﾝﾗｰﾒﾝ", 820},
    {"タンメン","ﾀﾝﾒﾝ", 750},
    {"天津丼","ﾃﾝｼﾝﾄﾞﾝ", 790},
    {"オムライス","ｵﾑﾗｲｽ", 700},


  }


  for _, food := range foods {
    log.Printf("%+v\n", food)
  }

  var foodb []Fline
  var temp1 Fline

  for i := 0; i < len(foods) ; i++ {
    temp1.Name = foods[i].Name
    temp1.Kana = foods[i].Kana
    temp1.Price = foods[i].Price
    temp1.Nameb = []byte(foods[i].Name)
    temp1.Kanab = []byte(foods[i].Kana)
    foodb = append(foodb, temp1)
  }

  foodb = insertsort(foodb)

  for _, food := range foodb {
    log.Printf("商品名：　%s, 価格（税込）: %d\n", food.Name, food.Price)
  }

}
