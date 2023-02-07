package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{11, 13, 15, 17}))
	fmt.Println(findMin([]int{5, 1, 2, 3, 4}))
	fmt.Println(findMin([]int{2, 1}))
	fmt.Println(findMin([]int{1, 2}))
	fmt.Println(findMin([]int{3, 1, 2}))
}

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		//fmt.Println(left, right)
		mid := (left + right) / 2

		if nums[mid] >= nums[left] {
			if nums[mid] > nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[mid] > nums[right] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}

	return nums[left]
}
