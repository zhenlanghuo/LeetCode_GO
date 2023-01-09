package main

import "fmt"

func main() {
	fmt.Println(numTrees(2))
}

func numTrees(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		sum := 0
		for j := 1; j <= i; j++ {
			sum += dp[j-1] * dp[i-j]
		}
		dp[i] = sum
	}
	fmt.Println(dp)

	return dp[n]
}
