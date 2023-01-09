package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(canPartitionKSubsets([]int{2, 2, 2, 2, 3, 4, 5}, 4))
}

func canPartitionKSubsets(nums []int, k int) bool {
	n := len(nums)
	memory := make([]bool, 1<<uint(n))
	for i := 0; i < len(memory); i++ {
		memory[i] = true
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k

	sort.Ints(nums)
	fmt.Println(target)
	fmt.Println(nums)
	//reverse(nums)

	var backtrack func(start, k, bucket, useFlag int) bool
	backtrack = func(start, k, bucket, useFlag int) bool {
		if k == 0 {
			return true
		}

		if bucket == target {
			memory[useFlag] = backtrack(0, k-1, 0, useFlag)
			if memory[useFlag] {
				return true
			}
		}

		if !memory[useFlag] {
			return memory[useFlag]
		}

		for i := start; i < n; i++ {
			//fmt.Println(i, start, k, bucket, useFlag)
			if useFlag&(1<<uint(i)) == 1<<uint(i) {
				continue
			}

			if bucket+nums[i] > target {
				return false
			}

			// 使用 nums[start]
			useFlag = useFlag | 1<<uint(i)
			bucket += nums[i]
			memory[useFlag] = backtrack(i+1, k, bucket, useFlag)
			if memory[useFlag] {
				return true
			}
			useFlag = useFlag ^ 1<<uint(i)
			bucket -= nums[i]
		}

		memory[useFlag] = false
		return false
	}

	return backtrack(0, k, 0, 0)
}

//func canPartitionKSubsets(nums []int, k int) bool {
//	n := len(nums)
//	memory := make([]bool, 1<<uint(n))
//	for i := 0; i < len(memory); i++ {
//		memory[i] = true
//	}
//	sum := 0
//	for i := 0; i < n; i++ {
//		sum += nums[i]
//	}
//	if sum%k != 0 {
//		return false
//	}
//	target := sum / k
//
//	sort.Ints(nums)
//	fmt.Println(target)
//	fmt.Println(nums)
//	//reverse(nums)
//
//	var backtrack func(start, k, bucket, useFlag int) bool
//	backtrack = func(start, k, bucket, useFlag int) bool {
//		if k == 0 {
//			return true
//		}
//
//		if !memory[useFlag] {
//			return memory[useFlag]
//		}
//
//		for i := start; i < n; i++ {
//			if useFlag&(1<<uint(i)) == 1<<uint(i) {
//				continue
//			}
//
//			// 不使用 nums[start]
//			memory[useFlag] = backtrack(i+1, k, bucket, useFlag)
//			if memory[useFlag] {
//				return true
//			}
//
//			if bucket+nums[i] > target {
//				memory[useFlag] = false
//				return false
//				//continue
//			}
//
//			// 使用 nums[start]
//			useFlag = useFlag | 1<<uint(i)
//			if bucket+nums[i] == target {
//				// 重置start、bucket
//				memory[useFlag] = backtrack(0, k-1, 0, useFlag)
//				if memory[useFlag] {
//					return true
//				}
//			} else {
//				memory[useFlag] = backtrack(i+1, k, bucket+nums[i], useFlag)
//				if memory[useFlag] {
//					return true
//				}
//			}
//			useFlag = useFlag ^ 1<<uint(i)
//		}
//
//		return false
//	}
//
//	return backtrack(0, k, 0, 0)
//}

//func reverse(nums []int) {
//	start, end := 0, len(nums)-1
//	for start < end {
//		nums[start], nums[end] = nums[end], nums[start]
//	}
//}
