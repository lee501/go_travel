package main 

import(
        "io"
        "os"
        "strings"
)

func main() {
  reader := io.TeeReader(strings.NewReader("go语言中文网"), os.Stdout)
  reader.Read(make([]byte, 20))
}
