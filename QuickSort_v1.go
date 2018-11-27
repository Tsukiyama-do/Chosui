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

func partition(a []Fline, fi int,fj int ) ([]Fline, int, int, int) {


//  log.Println(time.Now())           // => "2015-05-05 07:23:30.757800829 +0900 JST"
  var pivot Fline
  pivot = a[fi]        //  一番左の値をピボットにしている
  var i, j int = fi, fj
//  log.Printf("before partitioning : fi, fj, pivot.price : %d %d %d", i, j, pivot.Price)
  i++      //  ピボットの右隣を最初の値としている
  for ; i < j  ; {      //  左右重ならない限り　評価・入れ替えを実施

      for ; i <= fj ; i++  {      //  i がオーバーフローしない範囲で実施 左から１つずつ増やしている
        if a[i].Price >= pivot.Price {      //  ピボット以上なら停止
          break
        }
        if  i > j {      //  左右が超えたら停止
          break
        }
      }
      if i >= fj || i > j {       //  i がオーバーフローあるいは左右が超えたら停止
        break
      }

      for ; j >= fi ; j-- {   //  j がオーバーフローしない範囲で実施　右から１つずつ減らしている
        if a[j].Price < pivot.Price { //  ピボット以上なら停止
          break
        }
        if j <= fi || i > j {   //  j がオーバーフローあるいは左右が超えたら停止
          break
        }
      }
      if j <= fi || i > j {    //  j がオーバーフローあるいは左右が超えたら停止
        break
      }

      if i < j {   //  左右が重ならないならば、
        a[i], a[j] = a[j], a[i]   //  値の交換
        i++   //  ひとつ右隣に移動
        j--   //  ひとつ左隣に移動
      }
      if i > j || i > fj || j < fi {   //  i, j がオーバーフローあるいは左右が超えたら停止
        break
      }
  }

    a[fi], a[i-1] = a[i-1], a[fi]

//  log.Printf("after partitioning : i, j, pivot.price : %d %d %d , below is result:", i, j, pivot.Price)
/*
  for num, food := range a {
    log.Printf("i： %d,p: %d\n", num, food.Price)
  }
*/
//  log.Println(time.Now())           // => "2015-05-05 07:23:30.757800829 +0900 JST"
  return a, fi, i-1, fj


}

func Quick(a []Fline, i int, j int) []Fline {

    b, l, m, h := partition(a, i, j)

    if l < m {
      Quick(b, l, m)
    }
    if m+1 < h {
      Quick(b, m+1, h)
    }
    return b
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
    log.Printf("%v\n", food)
  }

  var foodb []Fline

  for i := 0; i < len(foods) ; i++ {
    var temp1 Fline
    temp1.Name = foods[i].Name
    temp1.Kana = foods[i].Kana
    temp1.Price = foods[i].Price
    temp1.Nameb = []byte(foods[i].Name)
    temp1.Kanab = []byte(foods[i].Kana)
    foodb = append(foodb, temp1)
  }

    log.Printf("starting time : %v\n",time.Now())           // => "2015-05-05 07:23:30.757800829 +0900 JST"

  foodb = Quick(foodb,0, len(foodb)-1)

    log.Printf("completed time : %v\n",time.Now())           // => "2015-05-05 07:23:30.757800829 +0900 JST"

  for _, food := range foodb {
    log.Printf("商品名：%s, 価格（税込）:%d\n", food.Name, food.Price)
  }

}
