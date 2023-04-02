package main

import "sort"

func main() {

}

func minNumber(nums1 []int, nums2 []int) int {
	m := make(map[int]struct{})

	nums1Min := nums1[0]
	nums2Min := nums2[0]

	for _, num := range nums1 {
		nums1Min = min(nums1Min, num)
		m[num] = struct{}{}
	}

	sort.Ints(nums2)

	for _, num := range nums2 {
		nums2Min = min(nums2Min, num)
		if _, ok := m[num]; ok {
			return num
		}
	}

	return min(nums1Min*10+nums2Min, nums2Min*10+nums1Min)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
