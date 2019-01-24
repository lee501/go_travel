package main

import(
  "fmt"
  "io"
  "strings"
)

func main() {
  content := "this is go limitreader test"
  reader := strings.NewReader(content)
  limitReader := &io.LimitedReader{R: reader, N: 8}
  for limitReader.N > 0 {
    tmp := make([]byte, 2)
    limitReader.Read(tmp)
    fmt.Printf("%s %v\n", tmp, limitReader.N)
  }
}
