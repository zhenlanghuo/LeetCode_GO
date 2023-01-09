package main

import (
	"container/heap"
	"sort"
)

type MedianFinder struct {
	queMin *minHeap
	queMax *minHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		queMin: &minHeap{},
		queMax: &minHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.queMin.IntSlice) == 0 {
		heap.Push(this.queMin, -num)
		return
	}

	queMinTop := this.queMin.IntSlice[0]
	//queMaxTop := this.queMax.IntSlice[0]

	if num <= -queMinTop {
		heap.Push(this.queMin, -num)
		if len(this.queMin.IntSlice)-len(this.queMax.IntSlice) == 2 {
			heap.Push(this.queMax, -heap.Pop(this.queMin).(int))
		}
	} else {
		heap.Push(this.queMax, num)
		if len(this.queMax.IntSlice)-len(this.queMin.IntSlice) == 1 {
			heap.Push(this.queMin, -heap.Pop(this.queMax).(int))
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.queMin.IntSlice) == len(this.queMax.IntSlice) {
		return (float64(-(this.queMin.IntSlice[0])) + float64(this.queMax.IntSlice[0])) / 2.0
	}
	return float64(-(this.queMin.IntSlice[0]))
}

type minHeap struct {
	sort.IntSlice
}

func (m *minHeap) Push(x interface{}) {
	m.IntSlice = append(m.IntSlice, x.(int))
}

func (m *minHeap) Pop() interface{} {
	length := len(m.IntSlice)
	result := m.IntSlice[length-1]
	m.IntSlice = m.IntSlice[:length-1]
	return result
}
