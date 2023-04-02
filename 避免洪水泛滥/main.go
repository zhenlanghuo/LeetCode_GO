package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	//fmt.Println(avoidFlood([]int{1, 2, 0, 1, 2}))
	//fmt.Println(avoidFlood([]int{1, 2, 0, 0, 1, 2}))
	fmt.Println(avoidFlood([]int{1, 2, 0, 2, 3, 0, 1}))
}

type MinHeap [][]int

func (m MinHeap) Len() int           { return len(m) }
func (m MinHeap) Less(i, j int) bool { return m[i][0] < m[j][0] }
func (m MinHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m *MinHeap) Push(x interface{}) {
	*m = append(*m, x.([]int))
}

func (m *MinHeap) Pop() interface{} {
	length := len(*m)
	result := (*m)[length-1]
	*m = (*m)[:length-1]
	return result
}

func avoidFlood(rains []int) []int {
	n := len(rains)
	ans := make([]int, n)
	for i, r := range rains {
		if r == 0 {
			ans[i] = 1
		} else {
			ans[i] = -1
		}
	}

	preRainIndex := make(map[int]int)
	nextRainIndex := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		nextRainIndex[i] = math.MaxInt64
		if rains[i] != 0 {
			if _, ok := preRainIndex[rains[i]]; ok {
				nextRainIndex[i] = preRainIndex[rains[i]]
			}
			preRainIndex[rains[i]] = i
		}
	}

	//fmt.Println(nextRainIndex)

	isFull := make(map[int]bool)
	drainQueue := make(MinHeap, 0)
	for i, r := range rains {
		fmt.Println(i, r, isFull, drainQueue)
		if r != 0 {
			if isFull[r] {
				return nil
			}
			heap.Push(&drainQueue, []int{nextRainIndex[i], r})
			isFull[r] = true
		} else {
			if len(drainQueue) != 0 {
				top := heap.Pop(&drainQueue).([]int)
				ans[i] = top[1]
				isFull[top[1]] = false
			}
		}
	}

	return ans
}
