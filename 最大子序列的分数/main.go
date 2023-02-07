package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxScore([]int{1, 3, 3, 2}, []int{2, 1, 3, 4}, 3))
	fmt.Println(maxScore([]int{4, 2, 3, 1, 1}, []int{7, 5, 10, 9, 6}, 1))
}

type MinHeap []int

func (m MinHeap) Len() int            { return len(m) }
func (m MinHeap) Swap(i, j int)       { m[i], m[j] = m[j], m[i] }
func (m MinHeap) Less(i, j int) bool  { return m[i] < m[j] }
func (m *MinHeap) Push(v interface{}) { *m = append(*m, v.(int)) }
func (m *MinHeap) Pop() interface{} {
	v := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return v
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	n := len(nums1)
	nums2Index := make([]int, n)
	for i := 0; i < n; i++ {
		nums2Index[i] = i
	}
	sort.Slice(nums2Index, func(i, j int) bool {
		return nums2[nums2Index[i]] > nums2[nums2Index[j]]
	})

	//fmt.Println(nums2Index)

	minHeap := make(MinHeap, 0, k)
	sum := int64(0)
	for i := 0; i < k; i++ {
		minHeap = append(minHeap, nums1[nums2Index[i]])
		sum += int64(nums1[nums2Index[i]])
	}
	heap.Init(&minHeap)

	result := sum * int64(nums2[nums2Index[k-1]])
	for i := k; i < n; i++ {
		if nums1[nums2Index[i]] > minHeap[0] {
			sum -= int64(minHeap[0])
			sum += int64(nums1[nums2Index[i]])
			heap.Pop(&minHeap)
			heap.Push(&minHeap, nums1[nums2Index[i]])
		}
		if result < sum*int64(nums2[nums2Index[i]]) {
			result = sum * int64(nums2[nums2Index[i]])
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
