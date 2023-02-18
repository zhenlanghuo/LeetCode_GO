package main

import "fmt"

func main() {
	fmt.Println(applyOperations([]int{1, 2, 2, 1, 1, 0}))
}

func applyOperations(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
	}

	fmt.Println(nums)
	//l, r := 0, n-1
	//for l < r {
	//	for l < n && nums[l] != 0 {
	//		l++
	//	}
	//
	//	for r >= 0 && nums[r] == 00 {
	//		r--
	//	}
	//
	//	if l < r {
	//		nums[l], nums[r] = nums[r], nums[l]
	//	}
	//}

	temp := make([]int, n)
	count := 0
	for _, num := range nums {
		if num != 0 {
			temp[count] = num
			count++
		}
	}

	return temp
}
