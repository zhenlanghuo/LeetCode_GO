package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(maxStarSum([]int{1, 2, 3, 4, 10, -10, -20}, [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}, {3, 5}, {3, 6}}, 2))
	fmt.Println(maxStarSum([]int{-5}, [][]int{}, 0))
	fmt.Println(maxStarSum([]int{1, -8, 0}, [][]int{{1, 0}, {2, 1}}, 2))
	fmt.Println(maxStarSum([]int{-1, 0}, [][]int{}, 1))
}

func maxStarSum(vals []int, edges [][]int, k int) int {
	edgeMap := make(map[int][]int)
	for i := 0; i < len(vals); i++ {
		edgeMap[i] = []int{}
	}
	for _, edge := range edges {
		edgeMap[edge[0]] = append(edgeMap[edge[0]], edge[1])
		edgeMap[edge[1]] = append(edgeMap[edge[1]], edge[0])
	}

	result := math.MinInt32
	for node, nodes := range edgeMap {
		temp := vals[node]
		sort.Slice(nodes, func(i, j int) bool {
			return vals[nodes[i]] > vals[nodes[j]]
		})
		for i := 0; i < k && i < len(nodes); i++ {
			temp = max(temp, temp+vals[nodes[i]])
		}
		result = max(result, temp)
	}

	//if result == math.MinInt32 {
	//	return 0
	//}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
