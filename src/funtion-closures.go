package main 

import "fmt"

func adder() func(int) int {
  sum := 0
  return func(x int) int {
    sum += x
    return sum
  }
}

func main() {
  pos := adder()
  for i := 1; i < 4; i++ {
    m := pos(i)
    fmt.Println(m)
  }
}
