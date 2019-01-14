package main

import (
  "log"
//  "fmt"
  "os"
  "bufio"
  "time"
  "strconv"
)

func init() {
    log.SetFlags(log.Lshortfile)
}

const BUFSIZE = 1024 // 読み込みバッファのサイズ
const readF = "BMText.txt"
const kwd = "Los Altos"

type chk_box struct{
  idx int
  rchar byte
  mint int

}

func chk_mv_len(d byte, e []byte, ip int ) int {

  var v_list chk_box
  var sla []chk_box
  var jj int = 0
  var rst int

  for i:=ip ; i >=0 ; i-- {    // 移動先のスライスを作成
    if i == ip  {
        v_list.idx = i
        v_list.rchar = e[i]
        v_list.mint = len(e)
    } else {
       v_list.idx = i
       v_list.rchar = e[i]
       v_list.mint = jj
    }
    sla = append(sla, v_list)
    jj++
  }

  for i:=0 ; i <= len(sla)-1 ; i++ {    //  移動先のスライスを少し修正
      tempc := sla[i].rchar
      tempm := sla[i].mint
      tempidx := sla[i].idx
      for j:=i ; j <= len(sla)-1 ; j++ {
        if tempidx == sla[j].idx {

        } else {
            if tempc == sla[j].rchar {  // 該当の同じ文字があったとき
              if tempm < sla[j].mint {  //　移動数が大きいのがあれば
                sla[j].mint = tempm     // 小さい方を設定
              }
            }
        }
      }

  }

  rst = len(e)
  for _, ss := range sla {
      if ss.rchar == d {
        rst = ss.mint     // 移動数
      }
  }
//  log.Printf("chk %d, %v, moved to %d ", ip, sla, rst)

  return rst

}


func BMSearch(a []byte, b []byte) *[]string {

    var okrst []string
    var imove int


    // 比較と文字列移動の制御と、
    for i:=0;  i <= len(a)-len(b) ;i++ {
//      log.Printf("Curren pos and its char is : %d, %s .", i, string(a[i]))
      for j:=len(b)-1 ; j >= 0 ; j-- {
//        jj = j              // used later

        if a[i+j] != b[j] {  // バイト単位で比較
          imove = chk_mv_len(a[i+j], b, j)
          break   // 比較してNGの場合は、ループを出る。
        }
        if j == 0 {  //
          okrst = append(okrst,string(b) + ":" + strconv.Itoa(i) )   // ヒットした箇所をsliceに追加
//          log.Printf("Hit here is %s, %s",string(b), strconv.Itoa(i))
        }
      }
      if (i + imove) < len(a) {
        i = i + imove - 1       // 移動する桁数
      }
    }
  return &okrst

}

func main() {

  var ptime1, ptime2 int
  var num int = 0
// 時間計測開始
  log.Printf("starting time : %v\n",time.Now())           // => "2015-05-05 07:23:30.757800829 +0900 JST"
  ptime1 = time.Now().Nanosecond()



// ファイルオープン処理
  file, err := os.Open("./" + readF)
  if err != nil {
      // Openエラー処理
      panic(err)
      log.Println("Error occured at opening a file.")

  }
  defer file.Close()


  log.Println("Start  \n")

// バッファーを定義して、ファイルを１行ずつ読む
  buf := make([]byte, BUFSIZE)
  scw := []byte(kwd)  // string -> byte slice変換
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {  // find a line

    buf = scanner.Bytes()    // get one line of text from the file.
    num = num+1
//      log.Println("\n")
        log.Printf("Line %d : %s",num, string(buf))

    rst := BMSearch(buf,scw)   // searh operation
      if len(*rst) == 0 {
        log.Println("No found by searching")
      }else {
        for idx, words := range *rst{
          log.Printf("Found word is  %d, %s", idx,  words)
        }
      }

  }



//  log.Printf("%v/n",scw)
//  log.Printf("%v/n",string(scw))


  log.Println("End  \n")

  // 時間計測終了
    log.Printf("completed time : %v\n",time.Now())           // => "2015-05-05 07:23:30.757800829 +0900 JST"
    ptime2 = time.Now().Nanosecond()

  // 処理時間を表示
  log.Printf("Elapse time(Nanosecond) : %d\n", ptime2 - ptime1 )   // Elapse time



}
