package main

import "fmt"

func main() {
	fmt.Println(maximumSubarraySum([]int{1, 5, 4, 2, 9, 9, 9}, 3))
	fmt.Println(maximumSubarraySum([]int{4, 4, 4}, 3))
}

func maximumSubarraySum(nums []int, k int) int64 {
	m := make(map[int]int)
	l, r := 0, 0
	m[nums[0]] = 1
	sum, ans := int64(nums[0]), int64(0)
	for l < len(nums) {
		for r+1 < len(nums) && r-l+2 <= k && m[nums[r+1]] == 0 {
			r++
			sum += int64(nums[r])
			m[nums[r]] = 1
		}

		if r-l+1 == k {
			ans = max(ans, sum)
		}

		sum -= int64(nums[l])
		m[nums[l]]--
		l++
	}

	return ans
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
