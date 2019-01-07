// 使用闭包来生成Even偶数序列
package main 

import (
  "fmt"
)

func GenerateEven() func() uint{
  i := uint(0)
  // 返回func()
  return func() (retVal uint) {
    retVal = i
    i += 2
    return
  }
}

func main() {
  nextEven := GenerateEven()
  fmt.Println(nextEven())
  fmt.Println(nextEven())
  fmt.Println(nextEven())
}
