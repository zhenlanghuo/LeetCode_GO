package main

import "fmt"

func main() {
	fmt.Println(countVowelStrings(1))
	fmt.Println(countVowelStrings(2))
	fmt.Println(countVowelStrings(3))
}

func countVowelStrings(n int) int {

	dp := make([]int, 5)
	for i := 0; i < 5; i++ {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < 5; j++ {
			dp[j] += dp[j-1]
		}
	}

	ans := 0
	for i := 0; i < 5; i++ {
		ans += dp[i]
	}
	return ans
}
