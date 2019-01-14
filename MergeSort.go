package main

import (
//  "fmt"
//  "sort"
  "log"
  "time"
  "math/rand"
  "strconv"
)

func init() {
    log.SetFlags(log.Lshortfile)
}

type FoodItem struct {
  No string
  Price int
}

type indexP struct {
  debits int
  debite int
  credits int
  credite int
}

func div_s(i_sl []indexP, i_s int, i_e int ) ([]indexP, int) {

  var i_wk int
  var idtemp indexP
  var i_flag int = 0  // 0 : not completed,   1:  completed   completed or not
  if (i_e - i_s ) == 0 {
    i_flag = 1
  }else if (i_e - i_s ) == 1 {
    idtemp.debits = i_s
    idtemp.debite = i_s
    idtemp.credits = i_e
    idtemp.credite = i_e
    i_sl = append(i_sl, idtemp)
    i_flag = 0
  } else if (i_e - i_s ) == 2 {
    idtemp.debits = i_s
    idtemp.debite = i_s
    idtemp.credits = i_s+1
    idtemp.credite = i_e
    i_sl = append(i_sl, idtemp)
    idtemp.debits = i_s+1
    idtemp.debite = i_s+1
    idtemp.credits = i_e
    idtemp.credite = i_e
    i_sl = append(i_sl, idtemp)
    i_flag = 1
  } else {
    i_wk = (i_s + i_e)/2     //  ２つに分割する際の、値をもってくる
    idtemp.debits = i_s
    idtemp.debite = i_wk
    idtemp.credits = i_wk+1
    idtemp.credite = i_e
    i_sl = append(i_sl, idtemp)
    i_flag = 0
    if i_s < i_wk {
      i_sl, i_flag = div_s(i_sl, i_s, i_wk)
    }
    if i_wk+1 < i_e {
      i_sl, i_flag = div_s(i_sl, i_wk+1, i_e)
    }
  }
    return i_sl, i_flag
}

func Msort_core(a []FoodItem, b []indexP) []FoodItem {

    var templen int
    var dbs, crs []FoodItem
    templen = len(b)

    for i := (templen-1) ; i >= 0 ;i-- {

      // 元のスライスから、スライスを２つ作成する。
      if b[i].debits != b[i].credite {      //  分割できなかったものは、除く
        for j := b[i].debits ; j <= b[i].debite; j++ {
          dbs = append(dbs, a[j])
        }
        for k := b[i].credits ; k <= b[i].credite; k++ {
          crs = append(crs, a[k])
        }

/*
        for idx, tb := range dbs {
          log.Printf("dbs(%d) is %v\n", idx, tb)
        }
        for idx, tc := range crs {
          log.Printf("crs(%d) is %v\n", idx, tc)
        }
*/

        // 元のスライスを修正し、つくった２つのスライスをpop する

        for l := b[i].debits ; l <= b[i].credite; l++ {
          if len(dbs) == 0 {
            a[l] = crs[0]
            if len(crs) > 1 {
              crs = append(crs[:0],crs[1:]...)
            } else {
              crs = nil
            }
          } else if len(crs) == 0 {
            a[l] = dbs[0]
            if len(dbs) > 1 {
              dbs = append(dbs[:0],dbs[1:]...)
            } else {
              dbs = nil
            }
          } else if dbs[0].Price < crs[0].Price {
            a[l] = dbs[0]
            if len(dbs) > 1 {
              dbs = append(dbs[:0],dbs[1:]...)
            } else {
              dbs = nil
            }
          } else {
            a[l] = crs[0]
            if len(crs) > 1 {
              crs = append(crs[:0],crs[1:]...)
            } else {
              crs = nil
            }
          }
        }

        dbs = nil
        crs = nil
/*
        log.Printf("C is %d", i)
        for inn, kk := range a {
            log.Printf("%d %d", inn, kk.Price)
        }
        log.Println("\n")
*/
      }
  }
  return a

}


func sort_check(foods []FoodItem) bool {
  log.Printf("Check sort.\n")
  var bnm int
  for ino, food := range foods {
    if ino == 0 {
      bnm = food.Price
    }else{
      if bnm > food.Price {
        log.Printf("Not sorted. %d, %d", ino, food.Price)
        return false
      }else{
        bnm = food.Price
      }
    }
  }
  log.Printf("Check sort ended.\n")
  return true
}



func MergeSort(a []FoodItem) []FoodItem {

  var in_p []indexP
  var templen,i_flag  int
  var b []FoodItem

  templen = len(a)

    // slice分割し、その情報をIndexPへ蓄える
    in_p, i_flag = div_s(in_p, 0, templen-1)

/*
    for idx, pos := range in_p {
      log.Printf("in_p(%d) : %v\n", idx, pos)
    }
*/
    if i_flag == 0 {
      log.Println("Hello")
    }
    // IndexPのそれぞれを、後ろから、チェックし、ソートする
      b = Msort_core(a, in_p)

    return b
}


func main() {

  var foods []FoodItem
  var ptime1, ptime2  int
  size := 10000
  for i := 0; i < size; i++ {
      var items FoodItem
      items.No = strconv.Itoa(i) + "Shohin"
      items.Price = rand.Intn(size-1)
      foods = append(foods, items)
  }

/*
  for _, food := range foods {
    log.Printf("%s,%d\n", food.No, food.Price)
  }
*/
    log.Printf("starting time : %v\n",time.Now())           // => "2015-05-05 07:23:30.757800829 +0900 JST"
    ptime1 = time.Now().Nanosecond()
    ptimes := time.Now()

  foods = MergeSort(foods)

    log.Printf("completed time : %v\n",time.Now())           // => "2015-05-05 07:23:30.757800829 +0900 JST"
    ptime2 = time.Now().Nanosecond()
    ptimee := time.Now()
/*
  for _, food := range foods {
    log.Printf("商品No：%s, 価格:%d\n", food.No, food.Price)
  }
*/

    if sort_check(foods) == false {   // call sort check program.
      log.Printf("Failure of sort process.")
    }else{
      log.Printf("Sort OK!")
    }


      log.Printf("Elapse time(Nanosecond) : %d - %d = %d\n", ptime2, ptime1, ptime2 - ptime1 )   // Elapse time
    //  log.Printf("Elapse time2 is %d\n", ptime2 )   // Elapse time

      difference := ptimee.Sub(ptimes)
      log.Printf("difference = %v\n", difference)



}
