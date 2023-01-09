package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 4, 5}
	fmt.Println(firstMissingPositive(nums))
	fmt.Println(nums)
}

func firstMissingPositive(nums []int) int {

	for i := 0; i < len(nums); i++ {
		for nums[i] <= len(nums) && nums[i] != i+1 && nums[i] > 0 && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return len(nums) + 1
}

//func firstMissingPositive(nums []int) int {
//
//	hasOne := false
//	for i:=0;i<len(nums);i++{
//		if nums[i] == 1 {
//			hasOne = true
//		}
//
//		if nums[i] <= 0 {
//			nums[i] = 1
//		}
//	}
//
//	if hasOne {
//		for i:=0;i<len(nums);i++{
//			index := int(math.Abs(float64(nums[i])) - 1)
//			if index < len(nums) && nums[index] > 0 {
//				nums[index] = nums[index] * -1
//			}
//		}
//
//		for i:=0;i<len(nums);i++{
//			if nums[i] > 0 {
//				return i+1
//			}
//		}
//
//		return len(nums) + 1
//	} else {
//		return 1
//	}
//
//
//}
