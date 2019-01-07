// 空接口，指定0个方法
package main

import "fmt"

type I interface{}

func main() {
  var i I
  describe(i)

  i = 42
  describe(i)

  i = "go"
  describe(i)
}

func describe(i I){
  fmt.Printf("(%v, %T)", i, i)
}
