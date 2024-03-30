package lrucache

import (
	"container/list"
)

type Item struct {
	Key int
	Value int
}

type lru struct {
	capacity int
	items map[int]*list.Element
	queue *list.List
}

func New(cap int) Cache {
	return &lru{
		capacity: cap,
		items: make(map[int]*list.Element),
		queue: list.New(),
	}
}

func (c *lru) purge(){
	if elem := c.queue.Back(); elem != nil{
		item := c.queue.Remove(elem).(*Item)
		delete(c.items, item.Key)
	}
}

func (c *lru) Set(key, value int){
	if c.capacity == 0{
		return
	}
	if elem, exist := c.items[key]; exist == true{
		c.queue.MoveToFront(elem)
		elem.Value.(*Item).Value = value
		return
	}

	if c.queue.Len() == c.capacity{
		c.purge()
	}
	item := &Item{
		Key: key,
		Value: value,
	}
	elem := c.queue.PushFront(item)
	c.items[item.Key] = elem
}

func (c *lru) Get(key int) (int, bool){
	elem, exist := c.items[key]
	if exist == false{
		return 0, false
	}
	c.queue.MoveToFront(elem)
	return elem.Value.(*Item).Value, true
}

func (c *lru) Clear(){
	c.items = make(map[int]*list.Element)
	c.queue.Init()
}

func (c *lru) Range(f func(key, value int) bool){
		for e := c.queue.Back(); e != nil; e = e.Prev() {
			item := e.Value.(*Item)
			if !f(item.Key, item.Value) {
				break
			}
		}
}