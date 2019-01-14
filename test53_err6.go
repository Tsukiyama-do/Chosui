package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)
//   https://qiita.com/nayuneko/items/3c0b3c0de9e8b27c9548


type DownloadError struct {
  StatusCode int
}
func (e *DownloadError) Error() string {
    return "httpd error!"
}

type SaveError struct {
  Filename, Message string
}
func (e *SaveError) Error() string {
    return e.Message
}

func main() {
  retry_count := 0
  for {
    err := DownloadToFile("http://example.com/hogehoge.html", "/tmp/hoge.html")
    if err == nil {
      fmt.Println("download success!")
      os.Exit(0)
    }
    // errインスタンスの型判別
    switch e := err.(type) {
      // ダウンロードエラー
      case *DownloadError:
        fmt.Printf("ダウンロードエラー %s [retry=%d, code=%d]\n", e.Error(), retry_count, e.StatusCode)
        retry_count++
        if retry_count > 3 {
          fmt.Printf("retry count over\n")
          os.Exit(1)
        }
      case *SaveError:
          fmt.Printf("保存エラー: %s [filename=%s]\n", e.Message, e.Filename)
          os.Exit(2)
      default:
        fmt.Printf("その他のエラー: %s\n", err)
        os.Exit(3)
    }
  }
}

func DownloadToFile(url, filename string) error {
  // ファイルダウンロード
  resp, err := http.Get(url)
  if err != nil {
    return err
  }
  defer resp.Body.Close()
  // 200以外のレスポンスはダウンロードエラーとする
  if resp.StatusCode != http.StatusOK {
    return &DownloadError{StatusCode: resp.StatusCode}
  }
  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return err
  }
  // ファイルへ保存
  err = ioutil.WriteFile(filename, data, os.ModePerm)
  if err != nil {
    return &SaveError{Message: err.Error(), Filename: filename}
  }
  return nil
}
