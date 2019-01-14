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

type Hsort struct{
  No string
  Price int
  A int
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

func CreateHTree(hs []Hsort) []Hsort {

  hs_max := len(hs) - 1
  var i_parent, t_Price int
  var t_No string
  var flag bool = true
  if hs_max > 0 {
    for ; flag == true ; {
      for i := hs_max ; i > 0 ; i-- {
        i_parent = hs[i].A
        if hs[i_parent].Price < hs[i].Price {   // 親と子の価格の比較

          t_Price = hs[i_parent].Price   //  一時退避
          t_No = hs[i_parent].No  //  一時退避
          hs[i_parent].Price = hs[i].Price   // 入替え
          hs[i_parent].No = hs[i].No   // 入替え
          hs[i].Price = t_Price   // 入替え
          hs[i].No = t_No   // 入替え
          flag = false    //  入替えがあったら設定
        }
      }
      if flag == false {   // 入替えをやったら、再度実行
        flag = true
        } else {    // 入替えが発生しなかったら　ヒープは完成
          break
        }
    }
  }

  return hs

}

func HTreePops(hs []Hsort, bs []FoodItem) ([]Hsort, []FoodItem) {

    var tempFI FoodItem
    var templen int

    templen = len(hs)

    tempFI.No = hs[0].No     //  最上位のものを退避
    tempFI.Price = hs[0].Price   //  最上位のものを退避

    bs = append(bs, tempFI)    // FoodItem のスライスに、最上位のものを追加
    if templen > 1 {    //  hs が１件だけになったときは、交換は行わない
      hs[0].No = hs[templen-1].No         //  最下位の値を最上位に移動
      hs[0].Price = hs[templen-1].Price  //  最下位の値を最上位に移動
    }
    hs = hs[:templen-1]   // hs から最下位をpopする
    return hs, bs
}

func RevSort(a []FoodItem)  []FoodItem {

    var revs []FoodItem  //  逆ソートされるslice
    var templen int

    templen = len(a)

    for i := templen-1 ; i >= 0 ; i-- {
      revs = append(revs, a[i])
    }

    return revs
}




func HeapSort(a []FoodItem) []FoodItem {

  var hs []Hsort         //  関数内のwork-slice
  var bsort []FoodItem   //  逆ソートされた slice
  var itemp int

  //　Stage 1.  work sliceを作成
  for idx, item := range a {
      var hstemp Hsort
      hstemp.No = item.No
      hstemp.Price = item.Price
      if idx > 0 {         //
        if idx%2 == 1  {
          itemp = idx/2
          hstemp.A = itemp    //  親ノードのidxの値
        } else if idx%2 == 0 {
            itemp = idx/2
            hstemp.A = itemp - 1   //  親ノードのidxの値
        }
      } else {
          hstemp.A = 0
      }

      hs = append(hs, hstemp)      //  work sliceに追加
  }

/*
    for idx, food := range hs {
      log.Printf("hsort idx:%d, No:%s, Price:%d, A:%d\n", idx, food.No, food.Price, food.A)
    }
*/


//　Stage 2.  work slice を使って逆ソートしたsliceを作成
  for ; len(hs) > 0 ; {
//  Heap Tree の整理
    hs = CreateHTree(hs)
//  Heap Tree の整理終了

/*
    for idx, food := range hs {
      log.Printf("hsort2 idx:%d, No:%s, Price:%d, A:%d\n", idx, food.No, food.Price, food.A)
    }
*/

//  最上位の値を別のsliceにコピーし、work-sliceの末端を、最上位に移動して、末端をpopする
        hs, bsort = HTreePops(hs,bsort)
//  終了

  }

  //　Stage 3.  逆ソートしたsliceに対して、逆に並べ替える
    bsort = RevSort(bsort)


    return bsort

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
     bsort = HeapSort(foods)      //  ヒープソートのプログラム

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



  log.Printf("Elapse time(Nanosecond) : %d\n", ptime2 - ptime1 )   // Elapse time
//  log.Printf("Elapse time2 is %d\n", ptime2 )   // Elapse time
  difference := ptimee.Sub(ptimes)
  log.Printf("difference = %v\n", difference)



}
