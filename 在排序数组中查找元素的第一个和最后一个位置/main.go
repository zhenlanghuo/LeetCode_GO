package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{1, 2, 3, 3, 3, 5, 5, 6, 7, 8}, 5))
}

func searchRange(nums []int, target int) []int {
	first, last := search(nums, target, 0, len(nums)-1)
	return []int{first, last}
}

func search(nums []int, target int, start, end int) (int, int) {
	if start > end {
		return -1, -1
	}

	if start == end && nums[start] == target {
		return start, end
	}

	mid := (start + end) / 2

	if nums[mid] > target {
		return search(nums, target, start, mid-1)
	} else if nums[mid] < target {
		return search(nums, target, mid+1, end)
	} else {
		first, leftLast := search(nums, target, start, mid)
		_, rightLast := search(nums, target, mid+1, end)
		last := leftLast
		if rightLast != -1 {
			last = rightLast
		}
		return first, last
	}
}
