package main

import "strconv"

func main() {

}

func findColumnWidth(grid [][]int) []int {
	n, m := len(grid), len(grid[0])
	ans := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans[i] = max(ans[i], len(strconv.FormatInt(int64(grid[j][i]), 10)))
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
