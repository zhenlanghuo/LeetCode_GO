package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(countWays([][]int{{6, 10}, {5, 15}}))
	fmt.Println(countWays([][]int{{34, 56}, {28, 29}, {12, 16}, {11, 48}, {28, 54}, {22, 55}, {28, 41}, {41, 44}}))
}

func countWays(ranges [][]int) int {

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	fmt.Println(ranges)
	l := make([][]int, 0)
	l = append(l, ranges[0])

	for i := 1; i < len(ranges); i++ {
		if ranges[i][0] > l[len(l)-1][1] {
			l = append(l, ranges[i])
		} else {
			l[len(l)-1][0] = min(l[len(l)-1][0], ranges[i][0])
			l[len(l)-1][1] = max(l[len(l)-1][1], ranges[i][1])
		}
	}

	n := len(l)
	fmt.Println(n)

	ans := 2
	for i := 2; i <= n; i++ {
		ans = (ans + ans) % (1e9 + 7)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
