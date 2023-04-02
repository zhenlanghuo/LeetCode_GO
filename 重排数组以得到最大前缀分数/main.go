package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxScore([]int{2, -1, 0, 1, -3, 3, -3}))
	fmt.Println(maxScore([]int{-2, -3, 0}))
}

func maxScore(nums []int) int {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	ans := 0
	sum := 0
	for _, num := range nums {
		sum += num
		if sum > 0 {
			ans++
		}
	}

	return ans
}
