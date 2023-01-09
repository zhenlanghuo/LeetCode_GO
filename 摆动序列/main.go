package main

import "fmt"

func main() {
	fmt.Println(wiggleMaxLength([]int{1,1}))
}

func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	up := make([]int, len(nums)+1)
	down := make([]int, len(nums)+1)
	up[1] = 1
	down[1] = 1

	for i := 2; i <= len(nums); i++ {
		if nums[i-1] > nums[i-2] {
			up[i] = max(up[i-1], down[i-1]+1)
		} else {
			up[i] = up[i-1]
		}

		if nums[i-1] < nums[i-2] {
			down[i] = max(down[i-1], up[i-1]+1)
		} else {
			down[i] = down[i-1]
		}
	}

	fmt.Println(up)
	fmt.Println(down)

	return max(up[len(nums)], down[len(nums)])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
