// 指针接受者
package main 

import (
  "fmt"
  "math"
)

type MyStruct struct{
  X, Y float64
}

func (v MyStruct) Abs() float64 {
  return math.Sqrt(v.X + v.Y)
}

func (v *MyStruct) Scale(f float64){
  v.X = v.X * f
  v.Y = v.Y * f
}

func main() {
  v := MyStruct{3,4}
  fmt.Println(v.Abs())
  v.Scale(10)
  fmt.Println(v)
  fmt.Println(v.Abs())
}
