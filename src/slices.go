// package main 

// import "fmt"

// func main() {
//   p := []int{2,3,4,5,6,8}
//   fmt.Println("p ==", p)

//   slice := p[:3]
//   slice[0] = 200
//   fmt.Println(slice[0])

//   for i:=0; i<len(p); i++ {
//     fmt.Printf("p[%d] == %d\n", i, p[i])
//   }
// }
package main 
import "fmt"

func modify_arr(array [5]int) {
    array[0]=10
    fmt.Println("In modify_arr,array values is:",array)

}
func main() {
    arr:=[5]int{1,2,3,4,5}
    modify_arr(arr)
    fmt.Println("In main,array values is:",arr)

}
