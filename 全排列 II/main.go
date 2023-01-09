package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{1, 2, 3}))
}

var (
	result [][]int
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	result = make([][]int, 0)
	n := len(nums)
	used := make([]bool, n)

	backtrack(nums, nil, used)
	return result
}
func backtrack(nums, temp []int, used []bool) {
	if len(temp) == len(nums) {
		dst := make([]int, len(temp))
		copy(dst, temp)
		result = append(result, dst)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		used[i] = true
		backtrack(nums, append(temp, nums[i]), used)
		used[i] = false
	}
}
