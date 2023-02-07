package main

import "fmt"

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
	fmt.Println(numSubarrayProductLessThanK([]int{1, 2, 3}, 0))
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	result := 0
	left := 0
	temp := 1
	for right, num := range nums {
		temp *= num
		for left < len(nums) && temp >= k {
			temp = temp / nums[left]
			left++
		}
		if temp < k && right >= left {
			result += right - left + 1
		}
	}

	return result
}
