package main

import (
  "log"
//  "fmt"
  "os"
  "time"
  "strconv"
)

func init() {
    log.SetFlags(log.Lshortfile)
}

const BUFSIZE = 1024 // 読み込みバッファのサイズ
const readF = "KMPText.txt"
const kwd = "sessions"

func kmpSearch(a []byte, b []byte) *[]string {

    var okrst []string
    var bf byte
    var ilast,jj int

    // １桁目と同一文字があれば、その最初の位置を探す
    bf = b[0]
    ilast = len(b)  //
    for ti:=1; ti < len(b) ; ti++ {
        if b[ti] == bf {
          ilast = ti
          break
        }
    }

    // 比較と文字列移動の制御と、
    for i:=0; i < len(a);i++{
      for j:=0 ; j < len(b) ; j++ {
        jj = j              // used later
        if a[i+j] == b[j]{  // バイト単位で比較

        }else{
          break   // 比較してNGの場合は、ループを出る。
        }
        if j == len(b)-1 {  //
          okrst = append(okrst,string(b) + ":" + strconv.Itoa(i))   // ヒットした箇所をsliceに追加
        }
      }
      if jj > ilast {  // 移動する桁数が、最初の文字が２回めでるところを超えてはいけない！
        i = i + ilast        // 移動する桁数
      }else{
        i = i + jj           // 移動する桁数
      }
    }
  return &okrst

}

func main() {

  var ptime1, ptime2 int

// 時間計測開始
  log.Printf("starting time : %v\n",time.Now())           // => "2015-05-05 07:23:30.757800829 +0900 JST"
  ptime1 = time.Now().Nanosecond()



// ファイル読み込み処理
  file, err := os.Open("./" + readF)
  if err != nil {
      // Openエラー処理
      log.Println("Error occured at opening a file.")

  }
  defer file.Close()


  log.Println("Start  \n")

  buf := make([]byte, BUFSIZE)
  scw := []byte(kwd)
  for {
      n, err := file.Read(buf)
      if n == 0 {
          break   // ゼロバイト時
      }
      if err != nil {
          // Readエラー処理
          log.Println("Error occurred at reading a file.")
          break
      }

//      log.Println("\n")
      log.Printf("%s",string(buf[:n]))
  }

//  log.Printf("%v/n",scw)
//  log.Printf("%v/n",string(scw))

  rst := kmpSearch(buf,scw)
    if len(*rst) == 0 {
      log.Println("No found.")
    }else {
      for idx, words := range *rst{
        log.Printf(" %d, %s", idx,  words)
      }
    }


  log.Println("End  \n")

  // 時間計測終了
    log.Printf("completed time : %v\n",time.Now())           // => "2015-05-05 07:23:30.757800829 +0900 JST"
    ptime2 = time.Now().Nanosecond()

  // 処理時間を表示
  log.Printf("Elapse time(Nanosecond) : %d\n", ptime2 - ptime1 )   // Elapse time



}
