package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimizeArrayValue([]int{3, 7, 1, 6}))
	fmt.Println(minimizeArrayValue([]int{13, 13, 20, 0, 8, 9, 9}))
}

func minimizeArrayValue(nums []int) int {

	maxV, minV := math.MinInt64, math.MaxInt64
	for i := 0; i < len(nums); i++ {
		maxV = max(maxV, nums[i])
		minV = min(minV, nums[i])
	}

	fmt.Println(minV, maxV)

	check := func(minV int) bool {
		temp := make([]int, len(nums))
		for i := 0; i < len(nums); i++ {
			temp[i] = nums[i]
		}

		for i := len(nums) - 1; i >= 1; i-- {
			if temp[i] > minV {
				temp[i-1] = temp[i-1] + temp[i] - minV
				//fmt.Println(minV, temp)
			}
		}

		//fmt.Println(minV, temp)
		return temp[0] <= minV
	}

	left := minV
	right := maxV
	for left <= right {
		//fmt.Println(left, right)
		mid := (left + right) / 2
		if check(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	//fmt.Println(left, right)
	return left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
