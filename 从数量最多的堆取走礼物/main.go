package main

import (
	"container/heap"
	"math"
)

func main() {

}

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *MaxHeap) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

func pickGifts(gifts []int, k int) int64 {

	sum := int64(0)
	maxHeap := make(MaxHeap, 0, len(gifts))
	for _, gift := range gifts {
		sum += int64(gift)
		maxHeap = append(maxHeap, gift)
	}

	heap.Init(&maxHeap)

	for i := 0; i < k; i++ {
		max := heap.Pop(&maxHeap).(int)
		left := (int)(math.Sqrt((float64)(max)))
		sum -= int64(max - left)
		heap.Push(&maxHeap, left)
	}

	return sum
}
