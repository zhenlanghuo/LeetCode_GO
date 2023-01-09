package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(nextGreaterElement(230241))
}

func nextGreaterElement(n int) int {

	nums := make([]int, 0)
	for n > 0 {
		nums = append(nums, n%10)
		n = n / 10
	}

	reverse(nums)

	//fmt.Println(nums)

	// 从后遍历, 第一个变小的数的下标
	leftIndex := len(nums) - 1
	for leftIndex-1 >= 0 && nums[leftIndex-1] >= nums[leftIndex] {
		leftIndex--
	}
	if leftIndex == 0 {
		return -1
	}
	leftIndex--

	//fmt.Println(leftIndex)

	// 从后遍历, 第一个比leftIndex下标所在位置的数要大的数的下标
	rightIndex := len(nums) - 1
	for nums[rightIndex] <= nums[leftIndex] {
		rightIndex--
	}

	//fmt.Println(rightIndex)

	nums[leftIndex], nums[rightIndex] = nums[rightIndex], nums[leftIndex]
	reverse(nums[leftIndex+1:])
	//fmt.Println(nums)
	result := 0
	for i := 0; i < len(nums); i++ {
		result = result*10 + nums[i]
	}

	if result > math.MaxInt32 {
		return -1
	}

	return result
}

func reverse(nums []int) {
	start := 0
	end := len(nums) - 1
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
