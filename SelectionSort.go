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


func SelectionSort(a []FoodItem) []FoodItem {

  var tempF,wklF FoodItem
  var tempi int

  for i:=0 ; i < len(a); i++ {
    tempF = a[i]     //  現在の最小値
    tempi = i
    for j:=1 ; j < (len(a)-i) ; j++{
      if tempF.Price > a[i+j].Price {
        tempF = a[i+j]   // 最小値の最新化
        tempi = i+j  // 最小値のあった項目
      }
    }

    wklF = a[i]     //  右端の値を退避
    a[i] = tempF    // 右端の値を最終の最小値にする
    a[tempi] = wklF   //  右端の値を返してあげる

  }

    return a

}

func main() {

  var foods,bsort []FoodItem
//  var hsort []Hsort
  var ptime1, ptime2  int
  size := 10000
  for i := 0; i < size; i++ {   // Creating random numbers
      var items FoodItem
      items.No = strconv.Itoa(i) + "Shohin"
      items.Price = rand.Intn(size-1)
      foods = append(foods, items)
  }

  log.Println("Before sort. \n")

/*
  for _, food := range foods {       //   ソート前のsliceを表示
    log.Printf("%s,%d\n", food.No, food.Price)
  }
*/
    log.Printf("starting time : %v\n",time.Now())           // => "2015-05-05 07:23:30.757800829 +0900 JST"
    ptime1 = time.Now().Nanosecond()
    ptimes := time.Now()
     bsort = SelectionSort(foods)      //  挿入ソートのプログラム

    log.Printf("completed time : %v\n",time.Now())           // => "2015-05-05 07:23:30.757800829 +0900 JST"
    ptime2 = time.Now().Nanosecond()
    ptimee := time.Now()

  log.Println("After sort. \n")

/*
  for _, food := range bsort {      //  sort後のsliceを表示
    log.Printf("商品No：%s, 価格:%d\n", food.No, food.Price)
  }
*/

    if sort_check(bsort) == false {   // call sort check program.
      log.Printf("Failure of sort process.")
    }else{
      log.Printf("Sort OK!")
    }

  log.Printf("Elapse time(Nanosecond) : %d - %d = %d\n", ptime2, ptime1, ptime2 - ptime1 )   // Elapse time
//  log.Printf("Elapse time2 is %d\n", ptime2 )   // Elapse time

  difference := ptimee.Sub(ptimes)
  log.Printf("difference = %v\n", difference)


}
