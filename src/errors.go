package main 

import (
  "fmt"
  "time"
)

type MyError struct{
  when time.Time 
  what string
}

func (e *MyError) Error() string {
  return fmt.Sprintf("at %v, %s", e.when, e.what)
}

func run() error {
  return &MyError{time.Now(), "don not work"}
}

func main() {
  err := run()
  if err != nil {
    fmt.Println(err)
  }
}
