package main

import "sort"

func main() {
	minCost([]int{1, 32, 53, 532}, []int{})
}

func minCost(basket1 []int, basket2 []int) int64 {
	n := len(basket1)
	count := make(map[int]int)
	mn := basket1[0]
	for i := 0; i < n; i++ {
		count[basket1[i]]++
		count[basket2[i]]--
		mn = min(mn, min(basket1[i], basket2[i]))
	}

	array := make([]int, 0)
	for k, v := range count {
		if v%2 != 0 {
			return -1
		}
		for i := 0; i < abs(v)/2; i++ {
			array = append(array, k)
		}
	}

	result := int64(0)
	sort.Ints(array)
	for i := 0; i < len(array)/2; i++ {
		result += int64(min(array[i], mn*2))
	}

	return result
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
