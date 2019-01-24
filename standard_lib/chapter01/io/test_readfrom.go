package main 

import (
  "fmt"
  "strings"
  "io"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error){
  p := make([]byte, num)
  n, err := reader.Read(p)
  if n > 0 {
    return p[:n], nil
  }
  return p, err
}

func main() {
  str := strings.NewReader("this is test go io reader")
  data, err := ReadFrom(str, 100)
  if (err == nil){
      fmt.Println(data)
  }
  a := "aaa"
  fmt.Println(a)

  b := []byte("go语言中文网")
  fmt.Println(b)
}
