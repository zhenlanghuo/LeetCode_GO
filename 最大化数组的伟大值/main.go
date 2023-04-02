package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximizeGreatness([]int{1, 3, 5, 2, 1, 3, 1}))
	fmt.Println(maximizeGreatness([]int{1, 2, 3, 4}))
}

func maximizeGreatness(nums []int) int {

	sort.Ints(nums)
	ans := 0

	maxIndex := len(nums) - 1
	nextMaxIndex := len(nums) - 2
	for maxIndex >= 0 && nextMaxIndex >= 0 {
		if nums[maxIndex] > nums[nextMaxIndex] {
			maxIndex--
			nextMaxIndex--
			ans++
		} else {
			nextMaxIndex--
		}
	}

	return ans
}
