package main

import "fmt"

func main() {
	fmt.Println(findUnsortedSubarray([]int{6,6,5,4,3,3,9}))
}

func findUnsortedSubarray(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	if nums[0] > nums[len(nums)-1] {
		return len(nums)
	}

	leftIndex := 0
	rightIndex := len(nums) - 1

	// 找出leftIndex，使得nums[0:leftIndex+1]是递增的
	for leftIndex+1 < len(nums) && nums[leftIndex] <= nums[leftIndex+1] {
		leftIndex++
	}
	if leftIndex == len(nums)-1 {
		return 0
	}

	// 找出rightIndex，使得nums[rightIndex:]是递增的
	for rightIndex-1 >= 0 && nums[rightIndex] >= nums[rightIndex-1] {
		rightIndex--
	}
	if rightIndex <= 0 {
		return len(nums)
	}

	fmt.Println(leftIndex, rightIndex)

	min := nums[leftIndex]
	max := nums[rightIndex]
	for i := leftIndex; i <= rightIndex; i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}

	fmt.Println(min, max)

	minLeftIndex := 0
	for minLeftIndex < len(nums) && nums[minLeftIndex] <= min {
		minLeftIndex++
	}

	maxRightIndex := len(nums) - 1
	for maxRightIndex >= 0 && nums[maxRightIndex] >= max {
		maxRightIndex--
	}

	fmt.Println(minLeftIndex, maxRightIndex)

	return maxRightIndex - minLeftIndex + 1
}
