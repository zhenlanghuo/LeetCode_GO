package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minCost([]int{20, 1, 15}, 5))
	fmt.Println(minCost([]int{1, 2, 3}, 4))

}

func minCost(nums []int, x int) int64 {

	mn := int64(math.MaxInt64)
	mnCount := make([]int, len(nums))
	copy(mnCount, nums)

	i := 0
	for i <= len(nums) {
		sum := int64(0)
		for _, v := range mnCount {
			sum += int64(v)
		}
		sum += int64(i * x)
		mn = min(mn, sum)
		nums = append(nums, nums[0])
		nums = nums[1:]
		for j, v := range nums {
			if mnCount[j] > v {
				mnCount[j] = v
			}
		}
		i++
	}

	return mn
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
