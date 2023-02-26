package main

import "sort"

func main() {

}

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	m := make(map[int]int)
	for _, num := range nums1 {
		m[num[0]] += num[1]
	}

	for _, num := range nums2 {
		m[num[0]] += num[1]
	}

	ans := make([][]int, 0)
	for k, v := range m {
		ans = append(ans, []int{k, v})
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i][0] < ans[j][0]
	})

	return ans
}
