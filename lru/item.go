package lru

import "time"

type Item struct {
	Value      string
	Expiration int64
}

func (i *Item) Expired() bool {
	if i.Expiration == 0 {
		return false
	}
	return time.Now().Unix() > i.Expiration
}

func (i *Item) Len() int {
	return len(i.Value)
}

func NewItem(data string, expireTime int64) *Item {
	var time_ int64
	if expireTime > 0 {
		time_ = time.Now().Unix() + expireTime
	} else {
		time_ = -1
	}
	return &Item{data, time_}

}
