package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(longestPath([]int{-1, 0, 0, 1, 1, 2}, "abacbe"))
	fmt.Println(longestPath([]int{-1, 0, 0, 0}, "aabc"))

}

type MinHeap []int

func (m MinHeap) Len() int            { return len(m) }
func (m MinHeap) Less(i, j int) bool  { return m[i] < m[j] }
func (m MinHeap) Swap(i, j int)       { m[i], m[j] = m[j], m[i] }
func (m *MinHeap) Push(v interface{}) { *m = append(*m, v.(int)) }
func (m *MinHeap) Pop() interface{} {
	top := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return top
}

func longestPath(parent []int, s string) int {
	n := len(parent)
	child := make([][]int, n)
	for c, p := range parent {
		if p == -1 {
			continue
		}
		child[p] = append(child[p], c)
	}

	result := 0
	// 返回 以根节点为起点的相邻节点无相同字符的最长路径长度
	var dfs func(root int) int
	dfs = func(root int) int {
		maxReturn := 1
		minHeap := (MinHeap)([]int{0, 0})
		heap.Init(&minHeap)
		for _, c := range child[root] {
			childMaxReturn := dfs(c)
			if s[c] != s[root] {
				maxReturn = max(maxReturn, childMaxReturn+1)
				if childMaxReturn > minHeap[0] {
					heap.Pop(&minHeap)
					heap.Push(&minHeap, childMaxReturn)
				}
			}
		}

		result = max(result, minHeap[0]+minHeap[1]+1)
		return maxReturn
	}
	dfs(0)

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
