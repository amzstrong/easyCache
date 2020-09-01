package test

import (
	"easyCache/lru"
	"fmt"
	"testing"
)


func TestLpush(t *testing.T)  {
	dl:=lru.NewDoubleLinkedList()

	dl.Lpush(lru.NewDoubleLinkedNode("a"))
	dl.Lpush(lru.NewDoubleLinkedNode("b"))
	dl.Lpush(lru.NewDoubleLinkedNode("c"))
	dl.Lpush(lru.NewDoubleLinkedNode("d"))
	//
	//fmt.Println(dl.GetTail().Data)
	//fmt.Println(dl.GetHead().Data)
	//fmt.Println(dl.GetLen())


	res,ok:=dl.Search("c")
	if ok{
		fmt.Println("search ok :",res.Data)
		//dl.Remove(res)
		dl.MoveToEnd(res)
	}else {
		fmt.Println("search failed")
	}

	fmt.Println("-----")

	h:=dl.Head
	for h!=nil{
		fmt.Println(h.Data)
		h=h.Next
	}
}

func TestCahce(T *testing.T)  {

	c := lru.NewCache(10)
	c.Set("a","1234567890")
	d,ok:=c.Get("a")
	fmt.Println(d,ok)
}