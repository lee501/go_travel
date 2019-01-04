// package main

// import "fmt"

// type Vertex struct {
//   x int
//   y int
// }

// func main() {
//   v := Vertex{1,2}
//   fmt.Println(v.x)

//   p := &v
//   fmt.Println(p)
// }

package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main() {
  v := Vertex{1, 2}
  p := &v
  p.X = 20
  fmt.Println(v)
}
