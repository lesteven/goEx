package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println(os.Args[0:])
    for i := 0; i < len(os.Args); i++ {
        fmt.Println(os.Args[i], i)
    }
}
