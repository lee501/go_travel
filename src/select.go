// package main

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     timeout := make(chan int, 1)
//     go func() {
//         time.Sleep(3e9) // sleep 3 seconds
//         timeout <- 0
//     }()
//     ch := make(chan int)
//     select {
//         case <-ch:
//         case <-timeout:
//             fmt.Println("timeout!")
//     }
// }
package main
import (
    "fmt"
    "time"
)
func main() {
    chann := make(chan int)
    go enqueue(chann)
    for {
        select {
        case v, ok := <-chann:
            if ok {
                fmt.Println(v)
            } else {
                fmt.Println("close")
                return
            }
        default:
            fmt.Println("waiting")
        }
    }
}
func enqueue(chann chan int) {
    time.Sleep(3 * time.Second)
    chann <- 1
    close(chann)
}
