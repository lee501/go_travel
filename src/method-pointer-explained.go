// 将Abs和Scale方法重写为函数
package main 

import (
  "fmt"
  "math"
)

type Menu struct{
  X, Y float64
}

func Abs(m Menu) float64{
  return math.Sqrt(m.X*m.X + m.Y*m.Y)
}

func Scale(m *Menu, f float64) {
  m.X = m.X * f
  m.Y = m.Y * f
}

func main() {
  v := Menu{3, 4}
  Scale(&v, 10)
  fmt.Println(v)
  fmt.Println(Abs(v))
}
