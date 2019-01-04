package main 

import "fmt"

func main() {
  i, j := 42, 2701

  p := &i
  fmt.Println(*p)
  *p = 21
  fmt.Println(i, j)

  p = &j         // point to j
  *p = *p / 37   // divide j through the pointer
  fmt.Println(j)
}
