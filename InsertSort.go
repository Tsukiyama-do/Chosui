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

type FoodItem struct{
  No string
  Price int
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



func InsertSort(a []FoodItem) []FoodItem {

  var templen int
  var tempF FoodItem
  var ipos int

  templen = len(a)

  for i := 0; i < templen ; i++ {  // sliceの先頭から最後まで調べる
    if i > 0 {
      tempF = a[i]    //  調べる対象を退避しておく
      ipos = i
      for j := 1 ; j <= i ; j++ {     //  後ろからひとつづつ下がる
        if a[i-j].Price >= tempF.Price {   // temp以上の場合
          a[i-j+1] = a[i-j]    // 一つ後ろにコピー
          ipos = i-j     //  tempFが入るべき場所
        }else {
          break
        }
      }
      a[ipos] = tempF   // 挿入すべきところに挿入
    }
  }

    return a

}

func main() {

  var foods,bsort []FoodItem
//  var hsort []Hsort
  var ptime1, ptime2  int
  size := 30
  for i := 0; i < size; i++ {   // Creating random numbers
      var items FoodItem
      items.No = strconv.Itoa(i) + "Shohin"
      items.Price = rand.Intn(size-1)
      foods = append(foods, items)
  }

  log.Println("Before sort. \n")


  for _, food := range foods {
    log.Printf("%s,%d\n", food.No, food.Price)
  }

    log.Printf("starting time : %v\n",time.Now())           // => "2015-05-05 07:23:30.757800829 +0900 JST"
    ptime1 = time.Now().Nanosecond()
     bsort = InsertSort(foods)      //  挿入ソートのプログラム

    log.Printf("completed time : %v\n",time.Now())           // => "2015-05-05 07:23:30.757800829 +0900 JST"
    ptime2 = time.Now().Nanosecond()

  log.Println("After sort. \n")


  for _, food := range bsort {
    log.Printf("商品No：%s, 価格:%d\n", food.No, food.Price)
  }


    if sort_check(bsort) == false {   // call sort check program.
      log.Printf("Failure of sort process.")
    }else{
      log.Printf("Sort OK!")
    }

  log.Printf("Elapse time(Nanosecond) : %d\n", ptime2 - ptime1 )   // Elapse time
//  log.Printf("Elapse time2 is %d\n", ptime2 )   // Elapse time

}
