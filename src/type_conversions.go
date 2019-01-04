package main 

import (
  "fmt"
  "math"
)

func main() {
  var i int
  j := i
  var x, y int = 3, 4
  var f float64 = math.Sqrt(float64(x*x + y*y))
  var z = int(f)
  fmt.Println(x, y, z, i, j)
}
