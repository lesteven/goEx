package main

import "fmt"

type Node struct {
    Val int
    Next *Node
}

func main() {
    var head *Node = &Node{Val: 10}
    var second *Node = &Node{Val: 55}
    head.Next = second
    second.Next = &Node{Val:100}

    for {
        if head == nil {
            break
        }
        fmt.Println(head.Val)
        head = head.Next
    }
}
