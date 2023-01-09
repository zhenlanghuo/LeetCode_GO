package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("ccc"))
}

func longestPalindrome(s string) string {
	n := len(s)

	if n <= 1 {
		return s
	}

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	result := string(s[0])

	for k := 1; k < n; k++ {
		for i := 0; i < n-k; i++ {
			j := i + k
			if s[i] != s[j] {
				continue
			}
			if i == j-1 {
				dp[i][j] = true
			} else {
				dp[i][j] = dp[i+1][j-1]
			}
			if dp[i][j] && len(result) < j-i+1 {
				result = s[i : j+1]
			}
		}
	}

	fmt.Println(dp)

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
