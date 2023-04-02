package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findScore([]int{2, 1, 3, 4, 5, 2}))
	fmt.Println(findScore([]int{2, 3, 5, 1, 3, 2}))
	fmt.Println(findScore([]int{2, 3, 5, 1, 6, 3, 2, 1}))
	//fmt.Println(findScore([]int{1}))
	//fmt.Println(findScore([]int{1, 2}))
	//fmt.Println(findScore([]int{1, 2, 3}))
	//fmt.Println(findScore([]int{3, 1, 2, 2}))
}

type Pair struct {
	index     int
	num       int
	heapIndex int
}

func findScore(nums []int) int64 {

	ans := int64(0)
	n := len(nums)
	pairs := make([]*Pair, n)
	for i, num := range nums {
		pairs[i] = &Pair{
			index: i,
			num:   num,
		}
	}

	visited := make([]bool, n)
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].num == pairs[j].num {
			return pairs[i].index < pairs[j].index
		}
		return pairs[i].num < pairs[j].num
	})

	for _, pair := range pairs {
		if visited[pair.index] {
			continue
		}
		visited[pair.index] = true
		ans += int64(pair.num)
		if pair.index-1 >= 0 {
			visited[pair.index-1] = true
		}
		if pair.index+1 < n {
			visited[pair.index+1] = true
		}
	}

	return ans
}

//func (p *Pair) less(a *Pair) bool {
//	if p.num == a.num {
//		return p.index < a.index
//	}
//	return p.num < a.num
//}
//
//func buildMinHead(arr []*Pair, headSize int) {
//	for i := headSize / 2; i >= 0; i-- {
//		minHeapify(arr, i, headSize)
//	}
//}
//
//func minHeapify(arr []*Pair, index, headSize int) {
//	left, right, largest := index*2+1, index*2+2, index
//	if left < headSize && arr[left].less(arr[largest]) {
//		largest = left
//	}
//	if right < headSize && arr[right].less(arr[largest]) {
//		largest = right
//	}
//	if largest != index {
//		arr[index], arr[largest] = arr[largest], arr[index]
//		arr[index].heapIndex = index
//		arr[largest].heapIndex = largest
//		minHeapify(arr, largest, headSize)
//	}
//}
//
//func findScore(nums []int) int64 {
//
//	n := len(nums)
//	pairs := make([]*Pair, n)
//	minHeap := make([]*Pair, n)
//	for i, num := range nums {
//		pair := &Pair{
//			index:     i,
//			num:       num,
//			heapIndex: i,
//		}
//		pairs[i] = pair
//		minHeap[i] = pair
//	}
//
//	buildMinHead(minHeap, n)
//
//	ans := int64(0)
//	heapSize := n
//
//	//fmt.Println(heapSize)
//
//	visited := make([]bool, n)
//	for len(minHeap) != 0 {
//		top := minHeap[0]
//		visited[top.index] = true
//		//fmt.Printf("top %v\n", top)
//		ans += int64(top.num)
//
//		minHeap[0], minHeap[heapSize-1] = minHeap[heapSize-1], minHeap[0]
//		minHeap[0].heapIndex = 0
//		minHeap = minHeap[:heapSize-1]
//		heapSize--
//		minHeapify(minHeap, 0, heapSize)
//
//		//fmt.Println(heapSize, ans)
//
//		if top.index-1 >= 0 && !visited[top.index-1] {
//			leftPair := pairs[top.index-1]
//			visited[top.index-1] = true
//			//fmt.Printf("leftPair %v\n", leftPair)
//			minHeap[leftPair.heapIndex], minHeap[heapSize-1] = minHeap[heapSize-1], minHeap[leftPair.heapIndex]
//			minHeap[leftPair.heapIndex].heapIndex = leftPair.heapIndex
//			minHeap = minHeap[:heapSize-1]
//			heapSize--
//			minHeapify(minHeap, leftPair.heapIndex, heapSize)
//		}
//
//		if top.index+1 < n && !visited[top.index+1] {
//			rightPair := pairs[top.index+1]
//			visited[top.index+1] = true
//			//fmt.Printf("rightPair %v\n", rightPair)
//
//			minHeap[rightPair.heapIndex], minHeap[heapSize-1] = minHeap[heapSize-1], minHeap[rightPair.heapIndex]
//			minHeap[rightPair.heapIndex].heapIndex = rightPair.heapIndex
//			minHeap = minHeap[:heapSize-1]
//			heapSize--
//			minHeapify(minHeap, rightPair.heapIndex, heapSize)
//		}
//	}
//
//	return ans
//}
