package main

import (
  "fmt"
  "github.com/syndtr/goleveldb/leveldb"

)


func main() {

// The returned DB instance is safe for concurrent use. Which mean that all
// DB's methods may be called concurrently from multiple goroutine.

  db, err := leveldb.OpenFile("/home/yuichi/github.com/Tsukiyama-do/Chosui/currentldb", nil)
  if err == nil {
  db.Put([]byte("k1"), []byte("Hello"), nil)
  db.Put([]byte("k2"), []byte("Good"), nil)
  db.Put([]byte("k3"), []byte("Morning"), nil)
  }

  iter := db.NewIterator(nil, nil)
  for iter.Next() {
  	// Remember that the contents of the returned slice should not be modified, and
  	// only valid until the next call to Next.
  	key := iter.Key()
  	value := iter.Value()
  	fmt.Printf("Key : %s, Value : %s \n", key , value)
  }
  iter.Release()
  err = iter.Error()

  db.Delete([]byte("k1"),nil)
  db.Delete([]byte("k2"),nil)
  db.Delete([]byte("k3"),nil)


  defer db.Close()

}
