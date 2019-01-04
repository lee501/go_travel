package main 

import "fmt"

func main() {
  var a [2]string
  a[0] = "hello"
  a[1] = "go"
  fmt.Println(a)
  fmt.Println(type(a[1]))
}
