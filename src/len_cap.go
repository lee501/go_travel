package main

import "fmt"

func main() {
  slice1 := []int{1,2,3,4,5}
  slice2 := make([]int, 5)
  slice3 := make([]int, 5, 6)
  slice4 := append(slice3, 6, 7, 8)

  print_info(slice1)
  print_info(slice2)
  print_info(slice3)
  print_info(slice4)
}

func print_info(slice []int) {
  fmt.Println("len: ", len(slice))
  fmt.Println("cap: ", cap(slice))

  for i, v := range slice {
    fmt.Println("element",i, ":", v)
  }
}
