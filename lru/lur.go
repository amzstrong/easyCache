package lru

import "time"

type Cache struct {
	maxBytes   int64
	usedBytes  int64
	linkedList *DoubleLinkedList
	items      map[string]*DoubleLinkedNode
	interval   time.Duration
}

func NewCache(maxBytes int64) *Cache {
	c := &Cache{
		maxBytes:   maxBytes,
		linkedList: NewDoubleLinkedList(),
		items:      make(map[string]*DoubleLinkedNode),
		interval:   10,
	}

	go c.gcLoop()

	return c
}

func (c *Cache) Get(key string) (value string, ok bool) {
	if ele, ok := c.items[key]; ok {
		c.linkedList.MoveToEnd(ele)
		return ele.Data.Value, true
	}
	return
}

func (c *Cache) Set(key string, data string, duration int64) bool {
	if ele, ok := c.items[key]; ok {
		ele.Data.Value = data
		ele.Data.Expiration = time.Now().Unix() + duration
		c.linkedList.MoveToEnd(ele)
		return true
	} else {
		newNode := NewDoubleLinkedNode(NewItem(data, duration))
		c.linkedList.Rpush(newNode)
		c.items[key] = newNode
		return true
	}
}

func (c *Cache) RemoveOld() {
	ele := c.linkedList.Head
	if ele != nil {
		c.linkedList.Remove(ele)
	}
}

func (c *Cache) Remove(key string) bool {
	if ele, ok := c.items[key]; ok {
		delete(c.items, key)
		c.linkedList.Remove(ele)
	}
	return true
}

func (c *Cache) gcLoop() {
	ticker := time.NewTicker(c.interval)
	for {
		select {
		case <-ticker.C:
			c.RemoveExpired()
		}
	}
}

func (c *Cache) RemoveExpired() {
	now := time.Now().Unix()
	for k, v := range c.items {
		if v.Data.Expiration > 0 && now > v.Data.Expiration {
			c.Remove(k)
		}
	}
}
