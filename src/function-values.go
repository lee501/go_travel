package main 

import (
  "fmt"
  "math"
)

func main() {
  test := func(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
  }

  fmt.Println(test(4,5))

  fmt.Println(compute(test))
}

func compute(fn func(float64, float64) float64) float64{
  return fn(4, 5)
}
