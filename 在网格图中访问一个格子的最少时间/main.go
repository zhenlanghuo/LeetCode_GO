package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumTime([][]int{{0, 1, 3, 2}, {5, 1, 2, 5}, {4, 3, 8, 6}}))
}

type MinHeap [][]int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i][2] < h[j][2] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(v interface{}) { *h = append(*h, v.([]int)) }
func (h *MinHeap) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

func minimumTime(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if grid[0][1] > 1 && grid[1][0] > 1 {
		return -1
	}

	ops := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	minHeap := make(MinHeap, 0)
	minHeap = append(minHeap, []int{0, 0, 0})
	dis := make([][]int, m)
	for i := 0; i < m; i++ {
		dis[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dis[i][j] = math.MaxInt64
		}
	}
	dis[0][0] = 0
	heap.Init(&minHeap)
	for {
		// item[0]为行, item[1]为列, item[2]是最小时间
		item := heap.Pop(&minHeap).([]int)
		//fmt.Println(item, minHeap)
		if item[0] == m-1 && item[1] == n-1 {
			return item[2]
		}
		for _, op := range ops {
			x, y, t := item[0]+op[0], item[1]+op[1], item[2]
			if x < 0 || x >= m || y < 0 || y >= n {
				continue
			}
			t = max(t+1, grid[x][y])
			if t%2 != (x+y)%2 {
				t += 1
			}
			if t < dis[x][y] {
				dis[x][y] = t
				heap.Push(&minHeap, []int{x, y, t})
			}

		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
