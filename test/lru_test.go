package test

import (
	"easyCache/lru"
	"fmt"
	"testing"
	"time"
)


func TestLpush(t *testing.T)  {
	dl:=lru.NewDoubleLinkedList()

	dl.Lpush(lru.NewDoubleLinkedNode(lru.NewItem("a",-1)))
	dl.Lpush(lru.NewDoubleLinkedNode(lru.NewItem("b",-1)))
	dl.Lpush(lru.NewDoubleLinkedNode(lru.NewItem("c",-1)))
	dl.Lpush(lru.NewDoubleLinkedNode(lru.NewItem("d",-1)))
	dl.Lpush(lru.NewDoubleLinkedNode(lru.NewItem("e",-1)))
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

	//c := lru.NewCache(10)
	//c.Set("a","1234567890")
	//d,ok:=c.Get("a")
	//fmt.Println(d,ok)
	a:=time.Now()
	fmt.Println(a)
}