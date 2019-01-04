package main

import "fmt"

func main() {
  a := make([]int, 5)
  printSlice("a", a)

  b := make([]int, 6)
  printSlice("b", b)

  d := b[2:5]
  printSlice("d", d)

  c := a[:2]
  printSlice("c", c)
  c[0] = 1
  fmt.Println(a)
}

func printSlice(s string, x []int) {
  fmt.Printf("%s len=%d cap=%d %v\n",
    s, len(x), cap(x), x)
}