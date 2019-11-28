package main

import "fmt"
import "io/ioutil"
import "syscall"

const con_path = "/home/yuichi/temp"

func main() {

  var i int = 0
  var s_st syscall.Stat_t

//    files, _ := ioutil.ReadDir("./")

    files, _ := ioutil.ReadDir(con_path)
    for _, f := range files {
//      if i > 1 {
        syscall.Stat(con_path + "/" + f.Name(), &s_st)

        fmt.Printf("%s  %s %s %10d %s %s \n",f.Mode(), s_st.Uid, s_st.Gid, f.Size(),   f.ModTime().Format("2006/01/02 15:04:05"),  f.Name())
//      }
      i++
    }
}
