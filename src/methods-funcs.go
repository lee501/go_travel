package main

import (
  "fmt"
  "math"
)

type T struct{
  X, Y float64
}

func Abs(v T) float64 {
  return math.Sqrt(v.X + v.Y)
}

func main() {
  v := T{3,4}
  fmt.Println(Abs(v))
}
