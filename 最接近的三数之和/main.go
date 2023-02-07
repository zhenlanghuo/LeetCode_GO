package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := math.MaxInt64
	for i, num := range nums {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			temp := num + nums[left] + nums[right]
			if target == temp {
				return target
			}
			if abs(temp-target) < abs(target-result) {
				result = temp
			}
			if temp > target {
				right--
			} else if temp < target {
				left++
			} else {
				return target
			}
		}
	}

	return result
}

func abs(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}
