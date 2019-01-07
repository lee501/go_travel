// go没有类，方法是具有接收者的函数
package main 

import (
  "fmt"
  "math"
)

type Stex struct {
  x, y float64
}

func (s Stex) abs() float64{
  return math.Sqrt(s.x*s.x + s.y*s.y)
}

func main() {
  v := Stex{5,6}
  fmt.Println(v.abs())
}
