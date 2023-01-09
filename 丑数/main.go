package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(nthUglyNumber(11))
}

func nthUglyNumber(n int) int {
	if n == 0 {
		return 0
	}

	pi := []int{0, 0, 0}
	iv := []int{2, 3, 5}
	dp := make([]int, n)
	dp[0] = 1

	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt32
		for j := 0; j < len(pi); j++ {
			dp[i] = min(dp[i], dp[pi[j]]*iv[j])
		}
		for j := 0; j < len(pi); j++ {
			if dp[i] == dp[pi[j]]*iv[j] {
				pi[j]++
			}
		}
		//fmt.Println(pi)
		//fmt.Println(dp)
	}

	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
