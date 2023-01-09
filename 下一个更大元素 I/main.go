package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElement([]int{2,4}, []int{1,2,3,4}))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nextGreater := make(map[int]int)
	stack := make([]int, 0)
	for _, num := range nums2 {
		for len(stack) != 0 {
			top := stack[len(stack)-1]
			if num > top {
				nextGreater[top] = num
				stack = stack[:len(stack)-1]
				continue
			}
			break
		}

		stack = append(stack, num)
	}

	for _, num := range stack {
		nextGreater[num] = -1
	}

	result := make([]int, 0, len(nums1))
	for _, num := range nums1 {
		result = append(result, nextGreater[num])
	}

	return result
}
