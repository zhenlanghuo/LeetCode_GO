package main

import "fmt"

func main() {
	fmt.Println(validSubarraySize([]int{1, 3, 4, 3, 1}, 6))
	fmt.Println(validSubarraySize([]int{6, 5, 6, 5, 8}, 7))
}

func validSubarraySize(nums []int, threshold int) int {

	n := len(nums)
	type Bound struct {
		left  int
		right int
	}
	bounds := make([]Bound, n)

	stack := make([]int, 0, n)
	for i := 0; i < n; i++ {
		for len(stack) != 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 {
			bounds[i].left = stack[len(stack)-1]
		} else {
			bounds[i].left = -1
		}
		stack = append(stack, i)
	}

	stack = stack[:0]
	for i := n - 1; i >= 0; i-- {
		for len(stack) != 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 {
			bounds[i].right = stack[len(stack)-1]
		} else {
			bounds[i].right = n
		}
		stack = append(stack, i)
	}

	fmt.Println(bounds)

	for i := 0; i < n; i++ {
		k := bounds[i].right - bounds[i].left - 1
		if float64(threshold)/float64(k) < float64(nums[i]) {
			return k
		}
	}

	return -1
}
