package main

import "fmt"

func main() {
	fmt.Println(findTheArrayConcVal([]int{7, 52, 2, 4}))
	fmt.Println(findTheArrayConcVal([]int{5, 14, 13, 8, 12}))
}

func findTheArrayConcVal(nums []int) int64 {
	result := int64(0)
	left, right := 0, len(nums)-1
	for left <= right {
		if left == right {
			result += int64(nums[left])
		} else {
			result += int64(nums[left]*pow10(numLen(nums[right]))) + (int64)(nums[right])
		}
		left++
		right--
	}
	return result
}

func numLen(num int) int {
	count := 0
	for num > 0 {
		num = num / 10
		count++
	}
	return count
}

func pow10(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}
