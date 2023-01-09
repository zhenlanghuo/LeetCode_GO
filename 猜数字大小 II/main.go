package main

import (
	"fmt"
)

func main() {
	fmt.Println(getMoneyAmount(17))
}

func getMoneyAmount(n int) int {
	if n <= 1 {
		return 0
	}
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := n - 1; i > 0; i-- {
		for j := i + 1; j <= n; j++ {
			dp[i][j] = j + dp[i][j-1]
			for k := i; k < j; k++ {
				cost := k + max(dp[i][k-1], dp[k+1][j])
				dp[i][j] = min(dp[i][j], cost)
			}
		}
	}

	fmt.Println(dp)

	return dp[1][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
