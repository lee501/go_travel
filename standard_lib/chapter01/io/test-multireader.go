package main 

import(
  // "fmt"
  // "strings"
  "io"
  // "bytes"
  "os"
)

func main() {
  // readers := []io.Reader{
  //   strings.NewReader("from strings reader"),
  //   bytes.NewBufferString("from bytes buffer"),
  // }

  // reader := io.MultiReader(readers...)
  // data := make([]byte, 0, 128)
  // buf := make([]byte, 10)

  // for n, err := reader.Read(buf); err != io.EOF ; n, err = reader.Read(buf){
  //   if err != nil{
  //       panic(err)
  //   }
  //   data = append(data,buf[:n]...)
  // }

  // fmt.Printf("%s\n", data)

  file, err := os.Create("tmp.txt")
  if err != nil {
      panic(err)
  }
  defer file.Close()
  writers := []io.Writer{
      file,
      os.Stdout,
  }
  writer := io.MultiWriter(writers...)
  writer.Write([]byte("Go语言中文网"))
}
