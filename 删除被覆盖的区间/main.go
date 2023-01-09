package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(removeCoveredIntervals([][]int{{1, 4}, {3, 3}, {2, 3}}))
}

func removeCoveredIntervals(intervals [][]int) int {

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	maxRight := intervals[0][1]
	count := 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i][1] > maxRight {
			count++
			maxRight = intervals[i][1]
		}
	}

	return count
}
