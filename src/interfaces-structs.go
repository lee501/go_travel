package main 

import "fmt"

type Phone interface {
  call()
  sales() int
}

type Nokia struct {
  price int
}

func (nokia Nokia) call() {
  fmt.Prinln()
}
