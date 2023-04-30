package main

func main() {

}

func findMaxFish(grid [][]int) int {

	ans := 0
	n, m := len(grid), len(grid[0])
	type Pair struct {
		row int
		col int
	}

	directs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		sum := 0
		if row < 0 || row >= n || col < 0 || col >= m {
			return 0
		}
		if grid[row][col] == 0 {
			return 0
		}
		sum += grid[row][col]
		grid[row][col] = 0
		for _, direct := range directs {
			sum += dfs(row+direct[0], col+direct[1])
		}
		return sum
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] > 0 {
				ans = max(ans, dfs(i, j))
			}
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
