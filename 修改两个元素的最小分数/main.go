package main

import "sort"

func main() {

}

func minimizeSum(nums []int) int {
	if len(nums) == 3 {
		return 0
	}

	sort.Ints(nums)
	ans := nums[len(nums)-1] - nums[0]
	ans = min(ans, nums[len(nums)-1]-nums[2])
	ans = min(ans, nums[len(nums)-2]-nums[1])
	ans = min(ans, nums[len(nums)-3]-nums[0])

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
