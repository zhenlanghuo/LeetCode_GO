package main

import "fmt"

func main() {
	fmt.Println(minInsertions("leetcode"))
}

func minInsertions(s string) int {

	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for k := 1; k < n; k++ {
		for i := 0; i < n-k; i++ {
			j := i + k
			dp[i][j] = min(dp[i+1][j]+1, dp[i][j-1]+1)
			if s[i] == s[j] {
				dp[i][j] = min(dp[i][j], dp[i+1][j-1])
			}
		}
	}

	return dp[0][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
