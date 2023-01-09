package main

import "fmt"

func main() {
	fmt.Println(combinationSum4([]int{9}, 3))
}

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			if i >= nums[j] {
				dp[i] += dp[i-nums[j]]
			}
		}
	}

	fmt.Println(dp)

	return dp[target]
}
