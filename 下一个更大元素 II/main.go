package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElements([]int{1, 5, 8, 2, 1}))
}

func nextGreaterElements(nums []int) []int {

	stack := make([]int, 0)
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = -1
	}

	for i := 0; i < 2*len(nums); i++ {
		index := i % len(nums)
		for len(stack) != 0 {
			top := stack[len(stack)-1]
			if nums[index] > nums[top] {
				result[top] = nums[index]
				stack = stack[:len(stack)-1]
				continue
			}
			break
		}
		stack = append(stack, index)
	}

	return result
}
