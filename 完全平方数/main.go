package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numSquares(10))
}

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minn := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minn = min(minn, dp[i-j*j])
		}
		dp[i] = minn + 1
	}

	fmt.Println(dp)

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
