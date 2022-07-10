package interfaces


type ILinkedList interface {
  Prepend(data interface{})
  Append(data interface{})
  Clear()
  Print()
}
