package interfaces

import (
  "github.com/braejan/practice-golang/src/linked-list/entities"
)

type ILinkedList interface {
  Prepend(node *entities.Node)
  Append(node *entities.Node)
  Clear()
  Print()
}
