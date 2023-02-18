package main

import "sort"

func main() {

}

func distinctAverages(nums []int) int {
	sort.Ints(nums)
	m := make(map[int]bool, 0)
	l, r := 0, len(nums)-1
	for l < r {
		m[nums[l]+nums[r]] = true
		l++
		r--
	}

	return len(m)
}
