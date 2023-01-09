package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(nthSuperUglyNumber(12, []int{2, 7, 13, 19}))
}

func nthSuperUglyNumber(n int, primes []int) int {
	// 记录质因数对应的指针
	indexList := make([]int, len(primes))
	nums := make([]int, len(primes))
	for i := 0; i < len(nums); i++ {
		nums[i] = primes[i]
	}
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt64
		for j := range indexList {
			dp[i] = min(dp[i], nums[j])
		}

		// 增加质因数的指针
		for j := range indexList {
			if dp[i] == nums[j] {
				indexList[j]++
				nums[j] = dp[indexList[j]] * primes[j]
			}
		}
	}

	//fmt.Println()

	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
