// 接口类型断言
package main 

import "fmt"

func main() {
  var i interface{}

  s, ok := i.(string)
  fmt.Println(s, ok)

  i = "go"
  s, ok = i.(string)
  fmt.Println(s, ok)

  f, ok := i.(float64)
  fmt.Println(f, ok)
}
