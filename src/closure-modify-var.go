// 使用闭包来修改局部变量
package main 

import(
  "fmt"
)

func main() {
  base := 100
  increase := func() {
    base += 100
  }
  fmt.Println(base)
  increase()
  fmt.Println(base)
}
