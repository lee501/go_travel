package main 

import "fmt"

func main() {
  var arr = []int{1,2,3,4}
  var base = 100
  // 可变参数 ...
  // 使用匿名函数作为闭包
  sum := func(base int, arr ...int) int {
    total := 0
    for _, val := range arr{
      total += val
    }
    return total + base
  }
  cal_sum := sum(base, arr...)
  fmt.Println(cal_sum)
}
