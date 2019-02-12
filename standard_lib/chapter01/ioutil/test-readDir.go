package main 

import(
  "fmt"
  "os"
  "ioutil"
)

func listAll(path string, cur int) {
  fileInfos, err := ioutil.ReadDir(path)
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, info := range fileInfos {
    if info.IsDir() {
      for tmpCur := cur; tmpCur > 0; tmpCur-- {
        fmt.Printf("|\t")
      }
      fmt.Println(info.Name(), "\\")
      listAll(path + "\\" + info.Name(),cur + 1)
    } else {
       for tmpHier := cur; tmpHier > 0; tmpHier--{
          fmt.Printf("|\t")
      }
      fmt.Println(info.Name())
    }
  }
}

func main() {
  dir := os.Args[0]
  fmt.Println(dir)
  listAll(dir, 0)
}
