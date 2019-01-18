// 信道
package main 

import "fmt"

func sum(s []int, c chan int){
  sum := 0
  for _, v := range s{
    sum += v
  }
  c <- sum
}

func main() {
  s := []int{7,2,8,9,-4,0}

  c := make(chan int)

  // 线程
  go sum(s[:], c)
  go sum(s[:len(s)/2], c)

  x, y := <-c, <-c
  fmt.Println(x, y)

  m := make(chan int)
  fmt.Println(<-m)
}

