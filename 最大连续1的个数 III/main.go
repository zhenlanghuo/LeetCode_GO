package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 0))
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 0))
}

func longestOnes(nums []int, k int) int {
	//zeroList := list.New()
	result := 0
	left := 0
	count := 0
	for right, num := range nums {
		if num == 0 {
			if k == 0 {
				left = right + 1
			} else {
				if count < k {
					count++
				} else {
					for nums[left] == 1 {
						left++
					}
					left++
				}
			}
		}

		if right >= left {
			result = max(result, right-left+1)
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
