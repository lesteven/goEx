package main

import "fmt"

func main() {
    val := 25
    addy := &val
    fmt.Printf("%v, %v, %v\n", val, addy, *addy)
    *addy = 100
    fmt.Printf("%v, %v, %v\n", val, addy, *addy)
}
