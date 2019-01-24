// IO包两个重要的接口：Reader和Writer
// os.File 同时实现了 io.Reader 和 io.Writer
// strings.Reader 实现了 io.Reader
// bufio.Reader/Writer 分别实现了 io.Reader 和 io.Writer
// bytes.Buffer 同时实现了 io.Reader 和 io.Writer
// bytes.Reader 实现了 io.Reader
// compress/gzip.Reader/Writer 分别实现了 io.Reader 和 io.Writer
// crypto/cipher.StreamReader/StreamWriter 分别实现了 io.Reader 和 io.Writer
// crypto/tls.Conn 同时实现了 io.Reader 和 io.Writer
// encoding/csv.Reader/Writer 分别实现了 io.Reader 和 io.Writer
// mime/multipart.Part 实现了 io.Reader
// net/conn 分别实现了 io.Reader 和 io.Writer(Conn接口定义了Read/Write)

// Reader接口的定义
// Read 将 len(p) 个字节读取到 p 中。它返回读取的字节数 n（0 <= n <= len(p)） 以及任何遇到的错误。
type Reader interface {
  Read(p []byte) (n int, err error)
}

// 所有实现了 Read 方法的类型都满足 io.Reader 接口
// 也就是说，在所有需要 io.Reader 的地方，可以传递实现了 Read() 方法的类型的实例
var EOF = errors.New("EOF")
// 具体使用该接口的用法
func ReadFrom(reader io.Reader, num int) ([]byte, error){
  p := make([]byte, num)
  n, err := reader.Read(p)
  if n > 0 {
    return p[:n], nil
  }
  return p, err
}
// 从标准输入读取
data, err = ReadFrom(os.Stdin, 11)

// 从普通文件读取，其中 file 是 os.File 的实例
data, err = ReadFrom(file, 9)

// 从字符串读取
data, err = ReadFrom(strings.NewReader("from string"), 12)



// ReaderAt和WriterAt接口
// ReaderAt 接口使得可以从指定偏移量处开始读取数据
type ReaderAt interface{
  ReadAt(p []byte, off int64) (n int, err error)
}

type WriterAt interface{
  WriteAt(p []byte, off int64) (n int, err error)
}

reader := strings.NewReader("go语言中文网")
p := make([]byte, 6)
n, err := reader.ReadAt(p, 2)
if err !=nil {
  panic(err)
}
fmt.Printf("%s, %d\n", p, n)

//  WriteAt 方法的使用（os.File 实现了 WriterAt 接口）
file, err := os.Create("writeAt.txt")
if err != nil {
  panic(err)
}
defer file.Close()
file.WriteString("Golang中文社区——这里是多余")
n, err := file.WriteAt([]byte("Go语言中文网")， 24)


// ReadFrom 从 r 中读取数据，直到 EOF 或发生错误。其返回值 n 为读取的字节数。除 io.EOF 之外，在读取过程中遇到的任何错误也将被返回。
type ReaderFrom interface{
  ReadFrom (r Reader) (n int64, err error)
}

file, err := os.Open("writeAt.txt")
if err != nil {
  panic(err)
}
defer file.Close()
writer := bufio.NewWriter(os.Stdout)
writer.ReadFrom(file)
writer.Flush()


// WriterTo的定义如下：

type WriterTo interface {
    WriteTo(w Writer) (n int64, err error)
}

reader := bytes.NewReader([]byte("Go语言中文网"))
reader.WriteTo(os.Stdout)


// closer接口： 文件 (os.File)、归档（压缩包）、数据库连接、Socket 等需要手动关闭的资源
type Closer interface {
    Close() error
}

func (f *File) Close() error {
    if f == nil {
        return ErrInvalid
    }
    return f.file.close()
}

// 1.8.1. ByteReader 和 ByteWriter
// bufio.Reader/Writer 分别实现了io.ByteReader 和 io.ByteWriter
// bytes.Buffer 同时实现了 io.ByteReader 和 io.ByteWriter
// bytes.Reader 实现了 io.ByteReader
// strings.Reader 实现了 io.ByteReader
type ByteReader interface {
    ReadByte() (c byte, err error)
}

type ByteWriter interface {
    WriteByte(c byte) error
}


var ch byte
fmt.Scanf("%c\n", &ch)

buffer := new(bytes.Buffer)
err := buffer.WriteByte(ch)
if err == nil {
    fmt.Println("写入一个字节成功！准备读取该字节……")
    newCh, _ := buffer.ReadByte()
    fmt.Printf("读取的字节：%c\n", newCh)
} else {
    fmt.Println("写入错误")
}

// ByteScanner、RuneReader 和 RuneScanner

type ByteScanner interface {
    ByteReader
    UnreadByte() error
}

buffer := bytes.NewBuffer([]byte{'a', 'b'})
buffer.readByte()
err := buffer.UnreadByte()

func Utf8Index(str, substr string) int {
  index := strings.Index(str, substr)
  if index < 0 {
    return -1
  }
  return utf8.RunCountInString(str[:index] )
}


// 若 dst 实现了 ReaderFrom 接口，其复制操作可通过调用 dst.ReadFrom(src) 实现。
// 此外，若 src 实现了 WriterTo 接口，其复制操作可通过调用 src.WriteTo(dst) 实现
func Copy(dst Writer, src Reader) (written int64, err error)
func CopyN(dst Writer, src Reader, n int64) (written int64, err error)

// ReadAtLeast 函数的签名：
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)


// MultiReader 和 MultiWriter 函数
func MultiReader(readers ...Reader) Reader
func MultiWriter(writers ...Writer) Writer


// TeeReader函数




























