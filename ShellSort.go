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

func findBand(a []FoodItem) []int {
  var templen,bn int
  var ii uint8
  var t_slice,b_slice []int

  //  バンドの計算を行っている。　バンドは、sliceの項目数を超えることはない。
  templen = len(a)
  if templen > 0 {
    for ii = 1 ;; ii++ {
      var ib2 int = ( 1 << ii ) - 1       //  2^ii -1
      if ib2 > templen {
         break
      }else {
        b_slice = append(b_slice, ib2)   //  バンドのsliceを作成　
      }
    }
  }

  bn = len(b_slice)
  if bn > 0 {
    for j := bn ; j > 0 ; j-- {
      t_slice = append(t_slice, b_slice[j-1])      //  バンドのsliceを順を逆にしている
    }
  }
/*
  for idx, val := range t_slice{
    log.Printf("find band t_slice [%d] : %d \n", idx, val)
  }
*/
    return t_slice
}



func bandSort(a []FoodItem, idx int, ib int) []FoodItem {

  var templen,j  int
  var tempF FoodItem

  //
  templen = len(a)
  for i := 1 ; (idx + ib*i) < templen ; i++ {
    j = i        //  バンド値に対して、sliceの項目の上限の項目値を洗い出し
  }

  // 上限項目まで、等差のsliceに対して、Insert Sort を行う
  for ; j > 0 ; j-- {
    for i := 1 ; i <= j ; i++ {
        if a[idx + ib*i].Price < a[idx + ib*(i-1)].Price {   // 右側が小さい場合は入替え
          tempF = a[idx + ib*(i-1)]
          a[idx + ib*(i-1)] = a[idx + ib*i]
          a[idx + ib*i] = tempF
      }
    }
  }
/*
  for idx, fd := range a{
      log.Printf("band sort  [%d] : %s  %d \n", idx, fd.No, fd.Price)
  }
*/
  return a

}




func ShellSort(a []FoodItem) []FoodItem {

  var iband []int

  // Step.1 Get band slice.

    iband = findBand(a)   // band のslice　を返す

  // 各　band に対して　Insert Sortを行う
    for _,ibve := range iband {      //  各バンドごとに実施
        for i := 0; i < ibve ; i++ {  // そのバンド内で　0〜　バンド−１　のそれぞれでソートさせる
          a = bandSort(a, i, ibve)    //  Insert Sort
        }
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
  for _, food := range foods {
    log.Printf("%s,%d\n", food.No, food.Price)
  }
*/
    log.Printf("starting time : %v\n",time.Now())           // => "2015-05-05 07:23:30.757800829 +0900 JST"
    ptime1 = time.Now().Nanosecond()
    ptimes := time.Now()
     bsort = ShellSort(foods)      //  挿入ソートのプログラム

    log.Printf("completed time : %v\n",time.Now())           // => "2015-05-05 07:23:30.757800829 +0900 JST"
    ptime2 = time.Now().Nanosecond()
    ptimee := time.Now()

  log.Println("After sort. \n")

/*
  for _, food := range bsort {
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
