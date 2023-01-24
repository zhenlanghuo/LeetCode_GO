package main

import "fmt"

func main() {
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
}

func minSubArrayLen(target int, nums []int) int {
	result := len(nums) + 1
	right := 0
	count := nums[0]
	for i := 0; i < len(nums); i++ {
		for right+1 < len(nums) {
			if count >= target {
				break
			}
			right = right + 1
			count += nums[right]
		}

		//fmt.Println(i, right, count)

		if count >= target && right-i+1 < result {
			result = right - i + 1
		}
		count -= nums[i]
	}

	if result == len(nums)+1 {
		return 0
	}

	return result
}
