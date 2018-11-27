package main

import (
  "fmt"
  "sort"
  "time"
)

type blockinfo struct{

  name string
  size uint16
  path string
  create_date string
  update_date string
}

func main() {

  var blockinfos = []blockinfo{
    {"Hanada", 20, "/tmp/te1", time.Now().String(),time.Now().String() },
    {"Narita", 25, "/tmp/te1", time.Now().String(),time.Now().String() },
    {"Asahikawa", 28, "/tmp/te1", time.Now().String(),time.Now().String() },
    {"Fukuoka", 29, "/tmp/te1", time.Now().String(),time.Now().String() },
    {"Chubu", 31, "/tmp/te1", time.Now().String(),time.Now().String() },
    {"Hong Kong", 33, "/tmp/te1", time.Now().String(),time.Now().String() },
    {"Kanazawa", 35, "/tmp/te1", time.Now().String(),time.Now().String() },
    {"Kochi", 28, "/tmp/te1", time.Now().String(),time.Now().String() },
  }

  var temp_size uint16
  temp_size = 28

  i := sort.Search(len(blockinfos), func(i int) bool {
        return blockinfos[i].size == temp_size
    })

    fmt.Printf("%v\n", blockinfos[i])

}
