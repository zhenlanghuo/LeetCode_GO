package main

import "fmt"

func main() {
	fmt.Println(integerBreak(10))
}

func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(dp[i-j], i-j)*max(dp[j], j))
		}
	}

	fmt.Println(dp)

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
