package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{7,7,7,7,7,7,7}))
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	max := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}
