package main

import "fmt"

type V struct {
  lat, long float64
}

var m = map[string]V{
  "lee": V{
    73, 175,
  },
}
func main() {
  fmt.Println(m)
}
