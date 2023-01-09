package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(triangleNumber([]int{0,0,0,1,1,2,3}))
}

//func triangleNumber(nums []int) (ans int) {
//	n := len(nums)
//	sort.Ints(nums)
//	for i, v := range nums {
//		k := i
//		for j := i + 1; j < n; j++ {
//			for k+1 < n && nums[k+1] < v+nums[j] {
//				k++
//			}
//			ans += max(k-j, 0)
//		}
//	}
//	return
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		k := i + 1
		for j := i + 1; j < len(nums)-1; j++ {
			for k < len(nums)-1 && nums[i]+nums[j] > nums[k+1] {
				k++
			}
			//if k-j > 0 {
			//	count += k - j
			//}
			count += k - j
		}
	}

	return count
}
