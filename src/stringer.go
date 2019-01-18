package main 

import "fmt"

type Person struct{
  name string
  age int
}

func (p *Person) String() string {
  return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

func main() {
  x := Person{name: "lee", age: 30}
  fmt.Println(&x)
}
