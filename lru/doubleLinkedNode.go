package lru

type DoubleLinkedNode struct {
	Data *Item
	Pre  *DoubleLinkedNode
	Next *DoubleLinkedNode
}

func NewDoubleLinkedNode(data *Item) *DoubleLinkedNode {
	return &DoubleLinkedNode{Data: data}
}
