package main

import "fmt"

func main() {
	fmt.Println(alternateDigitSum(521))
	fmt.Println(alternateDigitSum(886996))
}

func alternateDigitSum(n int) int {
	nums := make([]int, 0)
	for n > 0 {
		nums = append(nums, n%10)
		n = n / 10
	}
	one := 1
	result := 0
	for i := len(nums) - 1; i >= 0; i-- {
		result += nums[i] * one
		one *= -1
	}

	return result
}
