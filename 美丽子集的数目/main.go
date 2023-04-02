package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(beautifulSubsets([]int{2, 4, 6}, 2))
	fmt.Println(beautifulSubsets([]int{1}, 1))
}

func beautifulSubsets(nums []int, k int) int {

	sort.Ints(nums)
	n := len(nums)
	m := make(map[int]int)
	ans := 0
	count := 0
	var dfs func(start int)
	dfs = func(start int) {
		if start >= n {
			if count > 0 {
				ans++
			}
			return
		}

		// 不选
		dfs(start + 1)

		// 选
		if m[nums[start]-k] == 0 {
			m[nums[start]]++
			count++
			dfs(start + 1)
			m[nums[start]]--
			count--
		}
	}
	dfs(0)

	return ans
}
