package entities

import (
  "encoding/json"
)

type Node struct {
  Data interface{} `json:"data"`
  Next *Node `json:"data"`
}

func (n *Node) String() (toString string) {
  bytes, _ := json.Marshal(n)
  toString = string(bytes)
  return
}
