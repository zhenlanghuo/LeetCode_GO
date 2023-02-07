package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}

func findPeakElement(nums []int) int {

	nums = append(nums, math.MinInt64)
	nums = append([]int{math.MaxInt64}, nums...)

	left := 1
	right := len(nums) - 2
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > nums[mid+1] {
			// mid可能是峰顶或峰顶在mid的左侧
			right = mid - 1
		} else {
			// 峰顶在mid的右侧
			left = mid + 1
		}
	}

	return left - 1
}
