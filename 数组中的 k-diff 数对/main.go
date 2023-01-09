package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findPairs([]int{1,1,3,3,5}, 0))
}

func findPairs(nums []int, k int) int {
	if len(nums) <= 1 {
		return 0
	}

	sort.Ints(nums)

	lessIndex := 0
	moreIndex := 1

	count := 0

	for moreIndex < len(nums) {
		if nums[moreIndex]-nums[lessIndex] == k {
			if moreIndex == lessIndex {
				moreIndex++
				continue
			}
			count++
			for moreIndex+1 < len(nums) && nums[moreIndex] == nums[moreIndex+1] {
				moreIndex++
			}
			moreIndex++
			for lessIndex+1 < len(nums) && nums[lessIndex] == nums[lessIndex+1] {
				lessIndex++
			}
			lessIndex++
		} else if nums[moreIndex]-nums[lessIndex] > k {
			lessIndex++
		} else {
			moreIndex++
		}
	}

	return count
}
