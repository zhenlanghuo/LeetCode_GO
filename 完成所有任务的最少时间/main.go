package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMinimumTime([][]int{{2, 3, 1}, {4, 5, 1}, {1, 5, 2}}))
	fmt.Println(findMinimumTime([][]int{{1, 3, 2}, {2, 5, 3}, {5, 6, 2}}))
}

func findMinimumTime(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1] < tasks[j][1]
	})

	run := make([]bool, 2001)
	ans := 0
	for _, task := range tasks {
		start, end, d := task[0], task[1], task[2]
		for _, v := range run[start : end+1] {
			if v {
				d--
			}
		}
		if d <= 0 {
			continue
		}
		for i := end; d > 0; i-- {
			if !run[i] {
				run[i] = true
				d--
				ans++
			}
		}
	}

	return ans
}
