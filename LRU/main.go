// LRU
package main

import (
	"container/list"
	"fmt"
	"sync"
)

type LRUCache struct {
	mu       sync.Mutex
	list     *list.List
	entryMap map[int]*list.Element
	capacity int
}

type entry struct {
	key   int
	value int
}

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))
	lru.Put(3, 3)
	fmt.Println(lru.Get(2))
	lru.Put(4, 4)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(4))
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		list:     list.New(),
		entryMap: make(map[int]*list.Element),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	this.mu.Lock()
	defer this.mu.Unlock()

	el, ok := this.entryMap[key]
	if !ok {
		return -1
	}

	// 把元素放到队头
	this.list.MoveToFront(el)
	return el.Value.(*entry).value
}

func (this *LRUCache) Put(key int, value int) {
	this.mu.Lock()
	defer this.mu.Unlock()

	el, ok := this.entryMap[key]
	if !ok {
		// 新创建元素，放到队头
		el := this.list.PushFront(&entry{key: key, value: value})
		this.entryMap[key] = el
	} else {
		// 修改元素的value，放到队头
		el.Value.(*entry).value = value
		this.list.MoveToFront(el)
	}

	if this.list.Len() > this.capacity {
		// 超过容量，把队尾的元素淘汰
		el := this.list.Back()
		delete(this.entryMap, el.Value.(*entry).key)
		this.list.Remove(el)
	}
}
