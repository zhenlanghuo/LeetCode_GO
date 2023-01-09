package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{}))
}

func canJump(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	maxPosition := nums[0]
	for i := 1; i < n; i++ {
		if maxPosition >= i && maxPosition < i+nums[i] {
			maxPosition = i + nums[i]
		}
	}

	return maxPosition >= nums[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
