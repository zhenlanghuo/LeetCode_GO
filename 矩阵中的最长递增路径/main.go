package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(longestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}))
}

func longestIncreasingPath(matrix [][]int) int {

	rowLen, colLen := len(matrix), len(matrix[0])

	dp := make([][]int, rowLen)
	for i := 0; i < rowLen; i++ {
		dp[i] = make([]int, colLen)
		for j := 0; j < colLen; j++ {
			dp[i][j] = -1
		}
	}

	var memorySearch func(row, col int) int

	ops := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	result := 0
	memorySearch = func(row, col int) int {
		if dp[row][col] != -1 {
			return dp[row][col]
		}

		maxn := math.MinInt32
		for _, op := range ops {
			newRow, newCol := row+op[0], col+op[1]
			if newRow >= 0 && newRow < rowLen && newCol >= 0 && newCol < colLen {
				if matrix[row][col] > matrix[newRow][newCol] {
					maxn = max(maxn, memorySearch(newRow, newCol))
				}
			}
		}

		dp[row][col] = max(maxn, 0) + 1
		if result < dp[row][col] {
			result = dp[row][col]
		}
		return dp[row][col]
	}

	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if dp[i][j] == -1 {
				memorySearch(i, j)
			}
		}
	}

	//memorySearch(0, 0)

	fmt.Println(dp)

	return result
}



func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
