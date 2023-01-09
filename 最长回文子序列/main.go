package main

import "fmt"

func main() {
	fmt.Println(longestPalindromeSubseq("cbbd"))
}

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for k := 1; k < n; k++ {
		for i := 0; i < n-k; i++ {
			j := i + k
			dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			if s[i] == s[j] {
				dp[i][j] = max(dp[i][j], dp[i+1][j-1]+2)
			}
		}
	}

	//fmt.Println(dp)

	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
