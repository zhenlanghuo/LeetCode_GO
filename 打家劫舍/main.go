package main

import "fmt"

func main() {
	fmt.Println(rob([]int{}))
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = maxFunc(nums[0], nums[1])
	max := maxFunc(dp[0], dp[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = maxFunc(dp[i-2]+nums[i], dp[i-1])
		max = maxFunc(max, dp[i])
	}

	return max
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}
