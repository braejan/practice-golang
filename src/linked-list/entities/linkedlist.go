package entities

import (
  "fmt"
)

type LinkedList struct {
  Head *Node `json:"head"`
  Length int64 `json:"length"`
}

func toNode(data interface{}) (node *Node) {
  node = &Node{
    Data: data,
  }
  return
}

func NewLinkedList(args ...interface{}) (list *LinkedList){
  list = &LinkedList{}
  if args == nil {
    return
  } else {
    for _, v := range args {
      list.Append(v)
    }
  }
  return
}

func (list *LinkedList) Prepend (data interface{}) {
  if list.Head == nil {
    list.Head = toNode(data)
    list.Length++
    return
  }
  currentHead := list.Head
  list.Head = toNode(data)
  list.Head.Next = currentHead
  list.Length++
}

func (list *LinkedList) Append (data interface{}) {
  currentNode := list.Head
  if currentNode == nil {
    list.Head = toNode(data)
    list.Length++
    return
  }
  for currentNode.Next != nil {
    currentNode = currentNode.Next
  }
  currentNode.Next = toNode(data)
  list.Length++
}

func (list *LinkedList) Clear () {
  if list.Length == 0 {
    return
  }
  if list.Length == 1 {
    list.Length = 0
    list.Head = nil
    return
  }
  currentNode := list.Head
  for currentNode.Next != nil {
    aux := currentNode.Next
    currentNode.Next = nil
    currentNode = aux
  }
  list.Length = 0
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
