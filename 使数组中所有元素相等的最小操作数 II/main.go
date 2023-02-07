package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minOperations([]int{4, 3, 1, 4}, []int{1, 3, 7, 1}, 3))
	fmt.Println(minOperations([]int{3, 8, 5, 2}, []int{2, 4, 1, 6}, 1))
	fmt.Println(minOperations([]int{10, 5, 15, 20}, []int{20, 10, 15, 5}, 0))
}

func minOperations(nums1 []int, nums2 []int, k int) int64 {
	n := len(nums1)

	if k == 0 {
		for i := 0; i < n; i++ {
			if nums1[i] != nums2[i] {
				return -1
			}
		}
		return 0
	}

	diff := make([]int, n)
	diffSum := 0
	for i := 0; i < n; i++ {
		diff[i] = nums1[i] - nums2[i]
		if abs(diff[i])%k != 0 {
			return -1
		}
		diffSum += diff[i]
	}

	if diffSum != 0 {
		return -1
	}

	sort.Ints(diff)
	left := 0
	right := n - 1
	count := int64(0)
	for left < n && right >= 0 && left < right {
		for left < n && diff[left] == 0 {
			left++
		}

		for right >= 0 && diff[right] == 0 {
			right--
		}

		if left >= n || right < 0 || left >= right {
			break
		}

		times := min(abs(diff[left]), abs(diff[right])) / k
		diff[left] += times * k
		diff[right] -= times * k
		count += int64(times)
	}

	return count
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
