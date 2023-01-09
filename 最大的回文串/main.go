package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("aaccbbccab"))
}

func longestPalindrome(s string) string {

	result := ""

	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
		result = string(s[i])
	}

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			dp[i-1][i] = true
			result = s[i-1 : i+1]
		}
	}

	for k := 2; k < len(s); k++ {
		for i := 0; i+k < len(s); i++ {
			if s[i] == s[i+k] && dp[i+1][i+k-1] {
				dp[i][i+k] = true
				result = s[i : i+k+1]
			}
		}
	}

	return result
}
