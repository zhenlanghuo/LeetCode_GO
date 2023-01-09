package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {

	//fmt.Println(int(math.Ceil(float64(10) / float64(3))))

	fmt.Println(maxKelements([]int{1, 10, 3, 3, 3}, 3))
}

type maxHeap []int

func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Len() int           { return len(h) }
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() interface{} {
	top := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return top
}

func maxKelements(nums []int, k int) int64 {
	result := int64(0)
	h := maxHeap(nums)
	//fmt.Println(h)
	heap.Init(&h)

	//fmt.Println(h)

	for i := 0; i < k; i++ {
		//fmt.Println(h)
		top := heap.Pop(&h).(int)
		//fmt.Println(top)
		result += int64(top)
		//fmt.Println(int(math.Ceil(float64(top) / float64(3))))
		heap.Push(&h, int(math.Ceil(float64(top)/float64(3))))
	}

	return result
}
