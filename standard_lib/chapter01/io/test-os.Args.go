package main
import(
  "os"
  "fmt"
)
func main(){
  if len(os.Args)!=0{
    fmt.Println(os.Args[0])// args 第一个片 是文件路径
  }
  fmt.Println(os.Args[1])  // 第二个参数是， 用户输入的参数 例如 go run osdemo01.go 123
}
