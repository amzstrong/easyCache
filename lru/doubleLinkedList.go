package lru

type DoubleLinkedList struct {
	Head *DoubleLinkedNode
	Tail *DoubleLinkedNode
	Len  int
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

func (list *DoubleLinkedList) Rpush(node *DoubleLinkedNode) bool {
	if list.Len == 0 {
		list.Head = node
		list.Tail = node
	} else {
		n := list.Tail
		n.Next = node
		node.Pre = n
		list.Tail = node
	}
	list.Len = list.Len + 1
	return true
}

func (list *DoubleLinkedList) Lpush(node *DoubleLinkedNode) bool {
	if list.Len == 0 {
		list.Head = node
		list.Tail = node
	} else {
		n := list.Head
		node.Next = n
		n.Pre = node
		list.Head = node
	}
	list.Len = list.Len + 1
	return true
}

//lpop
func (list *DoubleLinkedList) Lpop(node *DoubleLinkedNode) *DoubleLinkedNode {

	if list.Len == 0 {
		return nil
	}
	old_head := list.Head
	next := list.Head.Next
	list.Head.Next = nil
	next.Pre = nil
	list.Head = next
	list.Len = list.Len - 1

	return old_head
}

//节点移动到末尾
func (list *DoubleLinkedList) MoveToEnd(node *DoubleLinkedNode) {
	list.Remove(node)
	if list.Len<=0{
		list.Head=node
		list.Tail=node
		list.Len=1
	}else{
		list.Tail.Next = node
		node.Pre = list.Tail
		node.Next = nil
		list.Tail = node
	}
}

func (list *DoubleLinkedList) Remove(node *DoubleLinkedNode) {
	p := list.Head
	if p == node && list.Len > 1 {
		list.Head = p.Next
		list.Len = list.Len - 1
		node=nil
	} else if p == node && list.Len == 1 {
		list.Head = nil
		list.Tail = nil
		list.Len = 0
		node=nil
	} else {
		for p.Next != nil {
			if p == node {
				ppre := p.Pre
				pnext := p.Next
				ppre.Next = pnext
				pnext.Pre = ppre
				node=nil
			}
			p = p.Next
		}
		list.Len = list.Len - 1
	}
}

func (list *DoubleLinkedList) Search(search string) (*DoubleLinkedNode, bool) {
	p := list.Head
	for p != nil {
		if p.Data == search {
			return p, true
		}
		p = p.Next
	}
	return nil, false
}
