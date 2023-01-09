// LFU
package main

import (
	"container/list"
	"fmt"
	"sync"
)

func main() {
	lfu := Constructor(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	fmt.Println(lfu.Get(1))
	lfu.Put(3, 3)
	fmt.Println(lfu.Get(2))
	fmt.Println(lfu.Get(3))
	lfu.Put(4, 4)
	fmt.Println(lfu.Get(1))
	fmt.Println(lfu.Get(3))
	fmt.Println(lfu.Get(4))
}

type KeyValue struct {
	key   int
	value int
	freq  int
}

type LFUCache struct {
	mu          sync.Mutex
	freqMap     map[int]*list.List
	keyValueMap map[int]*list.Element
	capacity    int
	minFreq     int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity:    capacity,
		minFreq:     1,
		freqMap:     make(map[int]*list.List),
		keyValueMap: make(map[int]*list.Element),
	}
}

func (this *LFUCache) Get(key int) int {
	this.mu.Lock()
	defer this.mu.Unlock()

	val, ok := this.keyValueMap[key]
	if !ok {
		return -1
	}

	keyValue := val.Value.(*KeyValue)

	this.incrFreq(val)

	return keyValue.value
}

func (this *LFUCache) Put(key int, value int) {
	this.mu.Lock()
	defer this.mu.Unlock()

	val, ok := this.keyValueMap[key]
	if ok {
		this.incrFreq(val)
		keyValue := val.Value.(*KeyValue)
		keyValue.value = value
		return
	}

	if _, ok := this.freqMap[1]; !ok {
		this.freqMap[1] = list.New()
	}
	val = this.freqMap[1].PushBack(&KeyValue{key: key, value: value, freq: 1})
	this.keyValueMap[key] = val

	if len(this.keyValueMap) > this.capacity {
		minFreqList, _ := this.freqMap[this.minFreq]
		if minFreqList != nil && minFreqList.Len() != 0 {
			minVal := minFreqList.Front()
			minFreqList.Remove(minVal)
			delete(this.keyValueMap, (minVal.Value.(*KeyValue)).key)
		}
	}
	this.minFreq = 1
}

func (this *LFUCache) incrFreq(val *list.Element) {
	keyValue := val.Value.(*KeyValue)

	prevFreq := keyValue.freq
	prevFreqList := this.freqMap[prevFreq]
	prevFreqList.Remove(val)

	nextFreq := prevFreq + 1
	if _, ok := this.freqMap[nextFreq]; !ok {
		this.freqMap[nextFreq] = list.New()
	}
	this.freqMap[nextFreq].PushBack(val.Value)

	// 更新 minFreq
	if this.freqMap[prevFreq].Len() == 0 && prevFreq == this.minFreq {
		this.minFreq = nextFreq
	}
}
