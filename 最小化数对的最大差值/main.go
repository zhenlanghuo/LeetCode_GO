package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimizeMax([]int{10, 1, 2, 7, 1, 3}, 2))
	fmt.Println(minimizeMax([]int{4, 2, 1, 2}, 1))
}

func minimizeMax(nums []int, p int) int {
	n := len(nums)
	sort.Ints(nums)
	check := func(diff int) bool {
		count := 0
		for i := 1; i < n; i++ {
			if nums[i]-nums[i-1] <= diff {
				count++
				i++
			}
		}
		if count >= p {
			return true
		}
		return count >= p
	}

	left, right := 0, int(1e9)
	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
