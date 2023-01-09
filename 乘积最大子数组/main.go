package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProduct([]int{-2, 2, 3, -2, 4}))
}

func maxProduct(nums []int) int {
	min, max, result := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		tempMax, tempMin := max, min
		//max = getMax(tempMax*nums[i], tempMin*nums[i], nums[i])
		//min = getMin(tempMax*nums[i], tempMin*nums[i], nums[i])
		//fmt.Println(max, min)

		if nums[i] > 0 {
			//max = maxFunc(tempMax*nums[i], nums[i])
			max = maxFunc(tempMax, 1) * nums[i]

			//min = minFunc(tempMin*nums[i], nums[i])
			min = minFunc(tempMin, 1) * nums[i]
		} else {
			//max = maxFunc(tempMin*nums[i], nums[i])
			max = minFunc(tempMin, 1) * nums[i]

			//min = minFunc(tempMax*nums[i], nums[i])
			min = maxFunc(tempMax, 1) * nums[i]
		}

		if max > result {
			result = max
		}
	}

	return result
}

func maxFunc(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func minFunc(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMax(a, b, c int) int {
	if c > a && c > b {
		return c
	}

	if b > a && b > c {
		return b
	}

	return a
}

func getMin(a, b, c int) int {
	if c < a && c < b {
		return c
	}

	if b < a && b < c {
		return b
	}

	return a
}
