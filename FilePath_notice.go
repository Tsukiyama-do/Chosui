package main
import(
  "os"
  "fmt"
  "log"
  "runtime"
  "path"
)

func main() {
  _, s_filename, _, ok := runtime.Caller(0)
  if !ok {
      panic("No caller information")
  }
  fmt.Printf("Filename : %q, Dir : %q\n", s_filename, path.Dir(s_filename))


  dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
  fmt.Println(dir)
}
