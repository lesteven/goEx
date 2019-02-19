package main

import "fmt"

func main() {
    arr := []int{1,2,5,10,20}
    fmt.Println(arr)

    // displays values
    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }
    fmt.Println()
    // displays values
    for _, each := range arr {
        fmt.Println(each)
    }
    fmt.Println()
    // displays keys
    for key := range arr {
        fmt.Println(key)
    }
    fmt.Println()
    // display keys and value
    for i, each := range arr {
        fmt.Println(i, each)
    }
}
