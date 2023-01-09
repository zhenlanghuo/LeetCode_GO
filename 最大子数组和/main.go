package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{5,4,-1,7,8}))
}

func maxSubArray(nums []int) int {
	max := nums[0]
	tempSum := 0

	for i := 0; i < len(nums); i++ {
		if tempSum < 0 && nums[i] > tempSum {
			tempSum = nums[i]
		} else {
			tempSum += nums[i]
		}

		if tempSum > max {
			max = tempSum
		}
	}

	return max
}
