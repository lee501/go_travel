package main 

import("fmt"
       "time"
      )

// func main() {
//   t := time.Now().Hour()
//   if t < 12 {
//     fmt.Println("morning")
//   } else if t < 18 {
//     fmt.Println("afternoon")
//   } else {
//     fmt.Println("evening")
//   }
// }

func main() {
  t := time.Now().Hour()
  switch {
  case t < 12:
    fmt.Println("morning")
  case t < 18:
    fmt.Println("afternoon")
  default:
    fmt.Println("evening")
  }
}
