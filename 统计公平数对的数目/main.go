package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(countFairPairs([]int{0, 1, 7, 4, 4, 5}, 3, 6))
	fmt.Println(countFairPairs([]int{1, 7, 9, 2, 5}, 11, 11))
}

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	//fmt.Println(nums)
	result := int64(0)
	for i, num := range nums {
		left, right := findFirstGETTarget(nums, i+1, lower-num), findLastLETTarget(nums, i+1, upper-num)
		//fmt.Println(i, num, left, right)
		if right >= left && left > 0 && right < len(nums) {
			result += int64(right - left + 1)
		}
	}

	return result
}

func findFirstGETTarget(nums []int, start, target int) int {
	left, right := start, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func findLastLETTarget(nums []int, start, target int) int {
	left, right := start, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left - 1
}
