package main

import (
    "io"
    "strings"
    "os"
)

func main() {
    r := strings.NewReader("some io.Reader stream to be read\n")
    lr := io.LimitReader(r, 4)
    io.Copy(os.Stdout, lr)
}
