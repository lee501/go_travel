// 指针作为接受者
// 方法能修改接受者指向的值
// 避免每次调用方法的时候赋值
package main

import (
  "fmt"
  "math"
)

type Vertux struct{
  X, Y float64
}

func (v *Vertux) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func (v *Vertux) Abs() float64{
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main() {
  v := &Vertux{ 3, 4 }
  fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
  v.Scale(5)
  fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
