package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxNumOfMarkedIndices([]int{3, 5, 2, 4}))
	fmt.Println(maxNumOfMarkedIndices([]int{9, 2, 5, 4}))
	fmt.Println(maxNumOfMarkedIndices([]int{42, 83, 48, 10, 24, 55, 9, 100, 10, 17, 17, 99, 51, 32, 16, 98, 99, 31, 28, 68, 71, 14, 64, 29, 15, 40}))
	fmt.Println(maxNumOfMarkedIndices([]int{7, 6, 8}))
}

func maxNumOfMarkedIndices(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	fmt.Println(nums)

	check := func(k int) bool {
		for i := 0; i < k; i++ {
			if 2*nums[i] > nums[n-k+i] {
				return false
			}
		}

		return true
	}

	l, r := 0, n/2
	for l <= r {
		mid := (l + r) / 2
		if check(mid) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return (l - 1) * 2
}
