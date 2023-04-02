package main

import "fmt"

func main() {
	fmt.Println(longestWPI([]int{9, 9, 6, 0, 6, 6, 9}))
	fmt.Println(longestWPI([]int{6, 6, 6, 9, 6, 6, 9, 9, 6, 6}))
}

func longestWPI(hours []int) int {

	n := len(hours)
	sums := make([]int, n+1)
	for i, hour := range hours {
		sums[i+1] = sums[i]
		if hour > 8 {
			sums[i+1]++
		} else {
			sums[i+1]--
		}
	}
	fmt.Println(sums)

	ans := 0
	m := make(map[int]int)
	for i, sum := range sums {
		if sum > 0 {
			//fmt.Println(i)
			ans = max(ans, i)
		} else {
			if v, ok := m[sum-1]; ok {
				ans = max(ans, i-v)
			}
		}
		if _, ok := m[sum]; !ok {
			m[sum] = i
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
