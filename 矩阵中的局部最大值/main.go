package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(largestLocal([][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}}))
	fmt.Println(largestLocal([][]int{{9, 9, 8}, {5, 6, 2}, {8, 2, 6}}))
	fmt.Println(largestLocal([][]int{{9, 9}, {5, 6}}))
}

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	if n <= 2 {
		return nil
	}
	result := make([][]int, n-2)
	for i := 0; i < n-2; i++ {
		result[i] = make([]int, n-2)
	}
	for i := 0; i+2 < n; i++ {
		for j := 0; j+2 < n; j++ {
			temp := math.MinInt32
			//fmt.Println(i, j)
			for i_ := i; i_ < i+3; i_++ {
				for j_ := j; j_ < j+3; j_++ {
					temp = max(temp, grid[i_][j_])
				}
			}
			result[i][j] = temp
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
