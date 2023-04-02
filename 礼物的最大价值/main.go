package main

import "fmt"

func main() {
	fmt.Println(maxValue([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}

func maxValue(grid [][]int) int {
	n, m := len(grid), len(grid[0])

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	dp[0][0] = grid[0][0]

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i-1 >= 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j]+grid[i][j])
			}

			if j-1 >= 0 {
				dp[i][j] = max(dp[i][j], dp[i][j-1]+grid[i][j])
			}
		}
	}

	//fmt.Println(dp)

	return dp[n-1][m-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
