package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestDivisibleSubset([]int{2, 3}))
}

func largestDivisibleSubset(nums []int) []int {

	sort.Ints(nums)

	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	maxSize := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if maxSize < dp[i] {
			maxSize = dp[i]
		}
	}

	fmt.Println(dp)

	result := make([]int, 0, maxSize)
	prev := 0
	for i := len(nums) - 1; i >= 0 && maxSize > 0; i-- {
		if dp[i] == maxSize && prev%nums[i] == 0 {
			result = append(result, nums[i])
			prev = nums[i]
			maxSize--
		}
	}

	reverse(result)

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverse(nums []int) {
	start, end := 0, len(nums)-1
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
