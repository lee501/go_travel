// ioutil库

// NopCloser函数
// 包装一个io.Reader,返回io.ReadCloser， 相应的CLose方法啥也不做，返回nil

// net/http包的NewRequest
// 接受一个io.Reader的body，Request的Body的类型是io.ReadCloser,
// 传递的io.Reader如果实现了io.ReadCloser接口，则转换否则通过ioutil.NopCloser包转换一下
rc, ok := body.(io.ReadCloser)
if !ok && body !=nil {
  rc = io.outil.NopCloser(body)
}

// ReadAll函数
// 一次性读取io.Reader中的数据
func ReadAll(r io.Reader) ([]byte, error)
// 源码通过bytes.Buffer中的ReadFrom中实现读取所有的数据，成功的时候返回的是err==nil
例：
var ch byte
fmt.Scanf("%c\n", &ch)

buffer := new(bytes.buffer)
err := buffer.WriteByte(ch)
if err == nil {
  fmt.Println("写入成功")
  newCh, _ := buffer.ReadByte()
  fmt.Printf("读取的字节：%c\n", newCh)
} else {
  fmt.Println("写入错误")
}



