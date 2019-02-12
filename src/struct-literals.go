package main

import "fmt"
import "os"

type Vertex struct{
  X, Y int
}

var (
  v1 = Vertex{1,2}
  v2 = Vertex{X: 1, Y: 3}
  v3 = Vertex{}
  p = &Vertex{1,2}
)

func main() {
  fmt.Println(v1.X, p.X, v2, v3)
  fmt.Println(os.Args[0])
}
