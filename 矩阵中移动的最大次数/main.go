package main

func main() {

}

func maxMoves(grid [][]int) int {

	n, m := len(grid), len(grid[0])
	memory := make([][]int, n)
	for i := 0; i < n; i++ {
		memory[i] = make([]int, m)
		for j := 0; j < m; j++ {
			memory[i][j] = -1
		}
	}

	directs := [][]int{{-1, 1}, {0, 1}, {1, 1}}

	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= n || y < 0 || y >= m {
			return -1
		}

		if memory[x][y] != -1 {
			return memory[x][y]
		}

		temp := -1
		for _, direct := range directs {
			nextX, nextY := x+direct[0], y+direct[1]
			if nextX < 0 || nextX >= n || nextY < 0 || nextY >= m {
				continue
			}
			if grid[x][y] >= grid[nextX][nextY] {
				continue
			}
			temp = max(temp, dfs(nextX, nextY))
		}
		memory[x][y] = temp + 1
		return temp + 1
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, dfs(i, 0))
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
