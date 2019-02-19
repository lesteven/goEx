package main

import "fmt"


func main() {
    foo := 10
    fmt.Println(foo)
    // doesnt overwrite value
    {
        foo := 200
        fmt.Println(foo)
    }
    fmt.Println(foo)

    // overwrites value
    {
        foo = 200
        fmt.Println(foo)
    }
    fmt.Println(foo)
}
