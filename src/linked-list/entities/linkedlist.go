package entities

import (
  "fmt"
)

type LinkedList struct {
  Head *Node `json:"head"`
  Length int64 `json:"length"`
}

func NewLinkedList(args ...interface{}) (list *LinkedList){
  list = &LinkedList{}
  if args == nil {
    return
  } else {
    for _, v := range args {
      node := &Node{
        Data: v,
      }
      list.Append(node)
    }
  }
  return
}

func (list *LinkedList) Prepend (node *Node) {
  if list.Head == nil {
    list.Head = node
    list.Length++
    return
  }
  currentHead := list.Head
  list.Head = node
  list.Head.Next = currentHead
  list.Length++
}

func (list *LinkedList) Append (node *Node) {
  currentNode := list.Head
  if currentNode == nil {
    list.Head = node
    list.Length++
    return
  }
  for currentNode.Next != nil {
    currentNode = currentNode.Next
  }
  currentNode.Next = node
  list.Length++
}

func (list *LinkedList) Clear () {
  list = &LinkedList{}
}

func (list *LinkedList) Print () {
  if list.Length == 0 {
    fmt.Printf("[]\n")
    return
  }
  if list.Length == 1 {
    fmt.Printf("[%v]\n", list.Head.Data)
    return
  }
  currentNode := list.Head
  fmt.Printf("[")
  for currentNode != nil {
    fmt.Printf("%v", currentNode.Data)
    currentNode = currentNode.Next
    if currentNode != nil {
      fmt.Printf(", ")
    }
  }
  fmt.Printf("]\n")
}
