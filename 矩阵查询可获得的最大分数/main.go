package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxPoints([][]int{{1, 2, 3}, {2, 5, 7}, {3, 5, 1}}, []int{5, 6, 2}))
}

type tuple struct {
	val int
	i   int
	j   int
}

type hp []tuple

func (h hp) Len() int               { return len(h) }
func (h hp) Swap(i, j int)          { h[i], h[j] = h[j], h[i] }
func (h hp) Less(i, j int) bool     { return h[i].val < h[j].val }
func (h *hp) Push(item interface{}) { *h = append(*h, item.(tuple)) }
func (h *hp) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func maxPoints(grid [][]int, queries []int) []int {
	n, m := len(grid), len(grid[0])
	ops := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	result := make([]int, len(queries))

	queryIndexes := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		queryIndexes[i] = i
	}
	//fmt.Println(queryIndexes)
	sort.Slice(queryIndexes, func(i, j int) bool {
		return queries[queryIndexes[i]] < queries[queryIndexes[j]]
	})
	//fmt.Println(queryIndexes)

	gridHeap := make(hp, 0, n*m)
	gridHeap = append(gridHeap, tuple{val: grid[0][0], i: 0, j: 0})
	heap.Init(&gridHeap)
	visited[0][0] = true

	count := 0
	for _, queryIndex := range queryIndexes {
		//fmt.Println(queryIndex, gridHeap)
		query := queries[queryIndex]
		for len(gridHeap) > 0 && gridHeap[0].val < query {
			count++
			cur := heap.Pop(&gridHeap).(tuple)
			for _, op := range ops {
				row := cur.i + op[0]
				col := cur.j + op[1]
				if row < 0 || row >= n || col < 0 || col >= m {
					continue
				}
				if visited[row][col] {
					continue
				}
				//if grid[row][col] >= query {
				//	continue
				//}
				heap.Push(&gridHeap, tuple{val: grid[row][col], i: row, j: col})
				visited[row][col] = true
			}
		}

		result[queryIndex] = count
	}

	return result
}
