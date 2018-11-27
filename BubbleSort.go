package main

import (
//  "fmt"
//  "sort"
  "log"
  "time"
  "math/rand"
)

func init() {
    log.SetFlags(log.Lshortfile)
}

type FoodItem struct {
  No int
  Price int
}


func partition(a []FoodItem, fi int,fj int ) ([]FoodItem, int, int, int) {


//  log.Println(time.Now())           // => "2015-05-05 07:23:30.757800829 +0900 JST"
  var pivot FoodItem
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
      if  i > j {       //  i がオーバーフローあるいは左右が超えたら停止
        break
      }

      for ; j >= fi ; j-- {   //  j がオーバーフローしない範囲で実施　右から１つずつ減らしている
        if a[j].Price < pivot.Price { //  ピボット以上なら停止
          break
        }
        if  i > j {   //  j がオーバーフローあるいは左右が超えたら停止
          break
        }
      }
      if  i > j {    //  j がオーバーフローあるいは左右が超えたら停止
        break
      }

      if i < j {   //  左右が重ならないならば、
        a[i], a[j] = a[j], a[i]   //  値の交換
        i++   //  ひとつ右隣に移動
        j--   //  ひとつ左隣に移動
      }
      if i == j {   //  万が一　上の交換で移動したあとに、同じ項目を参照していたら、
        if i <= fj {  //  オーバーフローチェック
          if a[i].Price < pivot.Price {   // 項目の値が、pivotより下だったら
            i++                           // iを右にひとつ移動
          }
        }
      }
      if i > j || i > fj || j < fi {   //  i, j がオーバーフローあるいは左右が超えたら停止
        break
      }
  }
    a[fi], a[i-1] = a[i-1], a[fi]

//  log.Printf("after partitioning : i, j, mid, pivot.price : %d %d %d, below is result:", i, j,  pivot.Price)
/*
  for num, food := range a {
    log.Printf("i： %d, %d\n", num, food.Price)
  }
*/
//  log.Println(time.Now())           // => "2015-05-05 07:23:30.757800829 +0900 JST"
  return a, fi, i-1, fj


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



func Bubble(a []FoodItem) []FoodItem {


  imax := len(a)-1
  vtemp := a[0]
  var flag bool
  for ; ; {
    flag = true
    for i := 1 ; i <= imax ; i++  {     //  右端に最大値をおく処理
      if a[i-1].Price > a[i].Price {
        vtemp = a[i]
        a[i] = a[i-1]
        a[i-1] = vtemp
        flag = false
      }
    }
    if flag == true {
      break
    }
  }
    return a
}


func main() {

  var foods []FoodItem
  var ptime1, ptime2  int
  size := 1000
  for i := 0; i < size; i++ {
      var items FoodItem
      items.No = i
      items.Price = rand.Intn(size-1)
      foods = append(foods, items)
  }

/*
  for _, food := range foods {
    log.Printf("%d,%d\n", food.No, food.Price)
  }
*/
    log.Printf("starting time : %v\n",time.Now())           // => "2015-05-05 07:23:30.757800829 +0900 JST"
    ptime1 = time.Now().Nanosecond()
  foods = Bubble(foods)

    log.Printf("completed time : %v\n",time.Now())           // => "2015-05-05 07:23:30.757800829 +0900 JST"
    ptime2 = time.Now().Nanosecond()
/*
  for _, food := range foods {
    log.Printf("商品No：%d, 価格:%d\n", food.No, food.Price)
  }
*/

    if sort_check(foods) == false {   // call sort check program.
      log.Printf("Failure of sort process.")
    }else{
      log.Printf("Sort OK!")
    }



  log.Printf("Elapse time(Nanosecond) : %d\n", ptime2 - ptime1 )   // Elapse time
//  log.Printf("Elapse time2 is %d\n", ptime2 )   // Elapse time
}
