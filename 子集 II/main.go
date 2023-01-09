package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2, 2}))
}

var (
	result [][]int
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result = make([][]int, 0)
	solve(nums, 0, make([]int, 0, 1<<uint64(len(nums))))
	return result
}

func solve(nums []int, start int, temp []int) {
	if start == len(nums) {
		//fmt.Println(temp)
		clone := make([]int, len(temp))
		copy(clone, temp)
		result = append(result, clone)
		return
	}

	end := len(nums)
	for i := start + 1; i < len(nums); i++ {
		if nums[i] != nums[start] {
			end = i
			break
		}
	}

	for i := start; i < end; i++ {
		temp = append(temp, nums[start])
		solve(nums, end, temp)
	}
	temp = temp[:len(temp)-(end-start)]
	solve(nums, end, temp)
}
