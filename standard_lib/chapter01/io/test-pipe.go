package main 

import(
  "fmt"
  "time"
  "io"
  "errors"
)

func main() {
  // io.Pipe() 用于创建一个同步的内存管道 (synchronous in-memory pipe)
  pipeReader, pipeWriter := io.Pipe()
  go PipeRead(pipeReader)
  go PipeWrit(pipeWriter)
  time.Sleep(30 * time.Second)
}

func PipeWrit(writer *io.PipeWriter) {
  data := []byte("go语言中文网")
  for i := 0; i < 3; i++ {
    n, err := writer.Write(data)
    if err != nil{
      fmt.Println(err)
      return
    }
    fmt.Printf("写入字节 %d\n", n)
  }
  writer.CloseWithError(errors.New("写入端关闭"))
}

func PipeRead(reader *io.PipeReader) {
  buf := make([]byte, 128)
  for {
    fmt.Println("接收端开始阻塞5秒钟...")
    time.Sleep(5 * time.Second)
    fmt.Println("接收端开始接受")
    n, err := reader.Read(buf)
    if err != nil {
      fmt.Println(err)
      return
    }
    fmt.Printf("收到字节：%d\n buf内容：%s\n", n, buf)
  }
}