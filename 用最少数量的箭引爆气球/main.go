package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
	fmt.Println(findMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}))
	fmt.Println(findMinArrowShots([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}))
}

func findMinArrowShots(points [][]int) int {

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	right := 1
	end := points[0][1]
	count := 1
	for right < len(points) {
		point := points[right]
		if point[0] > end {
			end = point[1]
			count++
		}
		right++
	}

	return count
}
