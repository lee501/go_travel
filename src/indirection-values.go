package main 

import (
  "fmt"
  "math"
)

type Vertex struct{
  X, Y float64
}

// 以值为接受者的方法， 接受者可以为值又能为指针
func (v Vertex) Abs() float64{
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
  v := Vertex{3, 4}
  p := &Vertex{3, 4}

// 接受者可以为值又能为指针
  fmt.Println(v.Abs())
  fmt.Println(p.Abs())

  fmt.Println(AbsFunc(v))
  fmt.Println(AbsFunc(*p))
}
