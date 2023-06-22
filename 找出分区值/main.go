package main

import (
	"math"
	"sort"
)

func main() {

}

func findValueOfPartition(nums []int) int {
	ans := math.MaxInt64
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		ans = min(ans, nums[i]-nums[i-1])
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
