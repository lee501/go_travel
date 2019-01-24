package main 

import(
      "fmt"
      // "bytes"
      "strings"
      )

// func main() {
//   var ch byte
//   fmt.Scanf("%c\n", &ch)

//   // n, err := fmt.Scanf("%s\n", &s)
//   // if err != nil {
//   //     fmt.Println("error", err)
//   // }

//   buffer := new(bytes.Buffer)
//   err := buffer.WriteByte(ch)
//   if err == nil {
//     fmt.Println("写入一个字节成功！准备读取该字节……")
//     newCh, _ := buffer.ReadByte()
//     fmt.Printf("读取的字节：%c\n", newCh)
//     } else {
//       fmt.Println("写入错误")
//     }
// }

func main() {
  // buffer := bytes.NewBuffer([]byte{'a', 'b'})
  // fmt.Printf("%c", buffer)
  var s = []string{'1', '2', 'c'} 
  index := strings.Index(s, 'c')
  fmt.Println(index)
}

