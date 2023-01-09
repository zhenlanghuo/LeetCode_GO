package main

import "fmt"

func main() {
	fmt.Println(search([]int{4,5,6,7,0,1,2}, 5))
}

func search(nums []int, target int) int {
	return searchWithIndex(nums, target, 0, len(nums)-1)
}

func searchWithIndex(nums []int, target, start, end int) int {
	if start > end {
		return -1
	}

	if start == end {
		if nums[start] == target {
			return start
		}
		return -1
	}

	if nums[start] == target {
		return start
	}

	if nums[end] == target {
		return end
	}

	if nums[start] > nums[end] {
		half := (start + end) / 2
		if nums[half] == target {
			return half
		}
		result := searchWithIndex(nums, target, start+1, half-1)
		if result != -1 {
			return result
		}
		return searchWithIndex(nums, target, half+1, end-1)
	}

	if nums[start] > target {
		return -1
	}

	if nums[end] < target {
		return -1
	}

	half := (start + end) / 2
	if nums[half] == target {
		return half
	}
	if nums[half] > target {
		return searchWithIndex(nums, target, start+1, half-1)
	}

	if nums[half] < target {
		return searchWithIndex(nums, target, half+1, end-1)
	}

	return -1
}
