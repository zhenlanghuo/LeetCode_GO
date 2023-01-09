package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(calculateMinimumHP([][]int{{0, -2, 3}, {-1, 0, 0}, {-3, -3, -2}}))
}

func calculateMinimumHP(dungeon [][]int) int {
	rowLen, colLen := len(dungeon), len(dungeon[0])
	dp := make([][]int, rowLen+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, colLen+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[rowLen-1][colLen] = 1
	dp[rowLen][colLen-1] = 1

	for i := rowLen - 1; i >= 0; i-- {
		for j := colLen - 1; j >= 0; j-- {
			dp[i][j] = max(min(dp[i+1][j], dp[i][j+1])-dungeon[i][j], 1)
		}
	}

	return dp[0][0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
