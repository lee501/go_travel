package main 

import "fmt"

type Vertex struct {
  Lat, Long float64
}

var m map[string]Vertex

func main() {
  m = make(map[string]Vertex)
  m["test"] = Vertex{1, -2}
  fmt.Println(m)
  fmt.Println(m["test"])
}
