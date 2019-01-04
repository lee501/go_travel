package main

import "fmt"

func main() {
  m := make(map[string]int)

  m["sex"] = 1

  fmt.Println("the val: ", m["sex"])

  v ,ok := m["lee"]
  fmt.Println("The value:", v, "Present?", ok)
}
