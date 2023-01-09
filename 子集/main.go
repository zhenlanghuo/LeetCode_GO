package main

import (
	"fmt"
)

var (
	result [][]int
)

func main() {
	fmt.Println(subsets([]int{1, 2, 3, 4, 5, 6}))
}

func subsets(nums []int) [][]int {
	result = make([][]int, 0)
	solve(nums, make([]int, 0, 1<<uint64(len(nums))))
	return result
}

func solve(nums []int, temp []int) {
	if len(nums) == 0 {
		clone := make([]int, len(temp))
		copy(clone, temp)
		result = append(result, clone)
		return
	}

	solve(nums[1:], append(temp, nums[0]))
	solve(nums[1:], temp)
}
