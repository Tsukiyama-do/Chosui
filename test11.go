package main

import (
  "fmt"
  "os"
)

func main(){


  fmt.Printf("Group id is %d\n", os.Getegid())
  fmt.Printf("Pagesize is %d\n", os.Getpagesize())
  fmt.Printf("Pid is %d\n", os.Getppid())
  hname, ok := os.Hostname()
  if ok == nil {
    fmt.Printf("Hostname is %s\n", hname)

  }
  fkx, ok := os.Create("tempFx")
  if ok == nil {
    fkx.Write([]byte("Noodle"))
    fkx.Close()
  }


}
