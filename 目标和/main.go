package main

import "fmt"

func main() {
	fmt.Println(findTargetSumWays([]int{1,1,1,1,1}, 4))
}

func findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}

	if target > sum || target < -sum {
		return 0
	}

	target = target + sum

	if target%2 == 1 {
		return 0
	}

	dp := make([]int, target/2+1)
	dp[0] = 1

	for i := 0; i < n; i++ {
		for j := target / 2; j-nums[i] >= 0; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}

	return dp[target/2]
}
