package main 

import(
        "fmt"
        // "io"
        // "os"
)

func main() {
  // io.Copy(os.Stdout, os.Stdin)
  // fmt.Println("Got EOF -- bye")
  var s string
  s = "abc"
  b := []byte(s)
  fmt.Printf("%v", b)
}
