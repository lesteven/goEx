package main

import "fmt"

type LinkedList struct {
    Val int
    Next *LinkedList
}
func (n *LinkedList) Add(Val int) int {
    var node *LinkedList = n
    for {
        if node.Next == nil {
            break
        }
        node = node.Next
    }
    node.Next = &LinkedList{Val: Val}
    return Val
}

func (n LinkedList) Print() {
    var node *LinkedList = &n
    for {
        if node == nil {
            break
        }
        fmt.Printf("%v ", node.Val)
        node = node.Next
    }
    fmt.Println()
}


func main() {
    var list LinkedList = LinkedList{Val: 2}
    list.Print()
    list.Add(10)
    list.Add(50)
    list.Print()
}
