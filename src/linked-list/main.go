package main

import (
  "fmt"
  "github.com/braejan/practice-golang/src/linked-list/entities"
  //"github.com/braejan/practice-golang/src/linked-list/interfaces"
)

func main() {
  fmt.Println("Running LinkedList tests")
  linkedList := entities.NewLinkedList("Braejan", "David", "Arias", 32, "Años")
  fmt.Printf("length: %d\n", linkedList.Length)
  linkedList.Print()
  linkedList = entities.NewLinkedList("Bruch")
  fmt.Printf("length: %d\n", linkedList.Length)
  linkedList.Print()
  linkedList.Clear()
  linkedList.Append("David")
  linkedList.Prepend("Braejan")
}


