package main

import (

   "log"
//   "hash"
   "github.com/Logibox/rollinghash"
)


func main(){

// s := []byte("The quick brown fox jumps over the lazy dog")
 s := []byte("The quick brown")

// This example works with adler32, but the api is identical for all
// other rolling checksums. Consult the documentation of the checksum
// of interest.
rh := rollinghash.AddBytes(s)

      rslt := rh.Sum32()
      log.Println(rslt)

}
