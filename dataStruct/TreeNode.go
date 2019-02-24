package main

import "fmt"


type Node struct {
    Val int
    Left *Node
    Right *Node
}

func print(node *Node) {
    if node == nil {
        return
    }
    fmt.Println(node.Val)
    if node.Left != nil {
        print(node.Left)
    }
    if node.Right != nil {
        print(node.Right)
    }
}

func main() {
    tree := &Node{Val:10}
    tree.Left = &Node{Val:5}
    tree.Right = &Node{Val:20}
    tree.Left.Left = &Node{Val:60}
    print(tree)
}
