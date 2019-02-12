package main

import (
     "fmt"
     "flag"
)

func main() {
   foo := flag.Bool("foo", false, "Foo")
   bar := flag.String("bar", "222", "Bar")

   flag.Parse()
   fmt.Print(*foo, *bar)
}
