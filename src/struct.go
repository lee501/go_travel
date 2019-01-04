package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main() {
  fmt.Println(Vertex{1, 2})
  v := Vertex{3, 4}
  v.X = 5
  fmt.Println(v)
  fmt.Println(v.X)
}
