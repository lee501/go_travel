// type关键字判断接口值的类型
package main 

import "fmt"

func do(i interface{}) {
  switch v := i.(type) {
  case int:
    fmt.Printf("Type is %T and Twice %v is %v\n", v, v, v*2)
  case string:
    fmt.Printf("%q is %v bytes long\n", v, len(v))
  default:
    fmt.Printf("can not check the type %T!\n", v)
  }
}

func main() {
  do(21)
  do("go")
  do(true)
}
