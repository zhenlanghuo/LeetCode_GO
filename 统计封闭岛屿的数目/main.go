package main

import "fmt"

func closedIsland(grid [][]int) int {
	n, m := len(grid), len(grid[0])

	queue := make([][]int, 0)

	for i := 0; i < n; i++ {
		if grid[i][0] == 0 {
			grid[i][0] = 2
			queue = append(queue, []int{i, 0})
		}
		if grid[i][m-1] == 0 {
			grid[i][m-1] = 2
			queue = append(queue, []int{i, m - 1})
		}
	}

	for i := 0; i < m; i++ {
		if grid[0][i] == 0 {
			grid[0][i] = 2
			queue = append(queue, []int{0, i})
		}
		if grid[n-1][i] == 0 {
			grid[n-1][i] = 2
			queue = append(queue, []int{n - 1, i})
		}
	}

	fmt.Println(grid)
	ops := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for i := 0; i < len(queue); i++ {
		for _, op := range ops {
			row := queue[i][0] + op[0]
			col := queue[i][1] + op[1]
			if !(row >= 0 && row < n && col >= 0 && col < m) {
				continue
			}
			if grid[row][col] != 0 {
				continue
			}
			queue = append(queue, []int{row, col})
			grid[row][col] = 2
		}
	}

	count := 0
	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			if grid[i][j] == 0 {
				count++
				setOne(grid, n, m, i, j, ops)
			}
		}
	}

	return count
}

func setOne(grid [][]int, n, m, row, col int, ops [][]int) {
	if row < 0 || row >= n || col < 0 || col >= m {
		return
	}
	if grid[row][col] != 0 {
		return
	}
	grid[row][col] = 1
	for _, op := range ops {
		setOne(grid, n, m, row+op[0], col+op[1], ops)
	}
}
