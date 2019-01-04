package main 

import "fmt"

func main() {
  var a []int
  printSlice("a", a)
 
  b := append(a, 0)
  printSlice("b", b)

  c := append(b, 1)
  printSlice("c", c)

  d := append(c, 2)
  printSlice("d", d)
}

func printSlice(s string, x []int) {
  fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
