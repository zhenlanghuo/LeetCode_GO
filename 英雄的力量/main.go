package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sumOfPower([]int{2, 1, 4}))
	fmt.Println(sumOfPower([]int{1, 1, 1}))
}

func sumOfPower(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	fmt.Println(nums)

	dp[0][0] = nums[0] * nums[0] * nums[0]
	dp[0][1] = nums[0]

	mod := int(1e9 + 7)

	for i := 1; i < n; i++ {
		dp[i][0] = (dp[i-1][0] + (dp[i-1][1]+nums[i])*nums[i]*nums[i]%mod) % mod
		dp[i][1] = (dp[i-1][1] + dp[i-1][1] + nums[i]) % mod
	}

	return dp[n-1][0]
}
