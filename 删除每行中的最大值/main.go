package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(deleteGreatestValue([][]int{{1, 2, 4}, {3, 3, 1}}))
	fmt.Println(deleteGreatestValue([][]int{{10}}))
	fmt.Println(deleteGreatestValue([][]int{}))
	fmt.Println(deleteGreatestValue([][]int{{}}))
}

func deleteGreatestValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	n, m := len(grid), len(grid[0])

	for i := 0; i < len(grid); i++ {
		sort.Ints(grid[i])
	}
	//fmt.Println(grid)

	result := 0
	for j := 0; j < m; j++ {
		temp := grid[0][j]
		for i := 0; i < n; i++ {
			temp = max(temp, grid[i][j])
		}
		result += temp
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
