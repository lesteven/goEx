package main

import "fmt"

type Node struct {
    Val int
    Next *Node
}

type LinkedList struct {
    Head *Node
}

func (n *LinkedList) Add(val int) {
    var node *Node = n.Head
    if node == nil {
        n.Head = &Node{Val : val}
        return
    }
    for node.Next != nil {
        node = node.Next
    }
    node.Next = &Node{Val : val}
}

func (n *LinkedList) Print() {
    var node *Node = n.Head
    if node == nil {
        fmt.Println("empty!")
        return
    }
    for node != nil {
        fmt.Printf("%v ", node.Val)
        node = node.Next
    }
    fmt.Println()
}

func main() {
    linkedList := LinkedList{}
    linkedList.Print()
    linkedList.Add(10)
    linkedList.Print()
    linkedList.Add(40)
    linkedList.Print()
}
