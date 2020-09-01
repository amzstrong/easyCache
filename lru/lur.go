package lru

type Cache struct {
	maxBytes   int64
	usedBytes  int64
	linkedList *DoubleLinkedList
	cache      map[string]*DoubleLinkedNode
}


func NewCache(maxBytes int64) *Cache {
	return &Cache{
		maxBytes:   maxBytes,
		linkedList: NewDoubleLinkedList(),
		cache:      make(map[string]*DoubleLinkedNode),
	}
}

func (c *Cache) Get(key string) (value string, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.linkedList.MoveToEnd(ele)
		return ele.Data, true
	}
	return
}

func (c *Cache) Set(key string, data string) bool {
	if ele, ok := c.cache[key]; ok {
		ele.Data = data
		c.linkedList.MoveToEnd(ele)
		return true
	} else {
		newNode := NewDoubleLinkedNode(data)
		c.linkedList.Rpush(newNode)
		c.cache[key] = newNode
		return true
	}
}

func (c *Cache) RemoveOld() {
	ele := c.linkedList.Head
	if ele != nil {
		c.linkedList.Remove(ele)
	}
}
