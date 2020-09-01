package lru

type DoubleLinkedNode struct {
	Data string
	Pre  *DoubleLinkedNode
	Next *DoubleLinkedNode
}

func NewDoubleLinkedNode(data string) *DoubleLinkedNode {
	return &DoubleLinkedNode{Data: data}
}
