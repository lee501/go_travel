package main 

import(
  "fmt"
  "strings"
)

func main() {
  // reader := strings.NewReader("Go语言中文网")
  // reader.Seek(-6, io.SeekEnd)
  // r, _, _ := reader.ReadRune()
  // fmt.Printf("%c\n", r)
  reader := strings.NewReader("go语言中文网")
  p := make([]byte, 6)
  n, err := reader.ReadAt(p, 2)
  if err !=nil {
    panic(err)
  }
  fmt.Printf("%s, %d\n", p, n)
}
