package main

import "fmt"

func main() {
	fmt.Println(isPossibleToCutPath([][]int{{1, 1, 1, 0, 0}, {1, 0, 1, 0, 0}, {1, 1, 1, 1, 1}, {0, 0, 1, 1, 1}, {0, 0, 1, 1, 1}}))
}

func isPossibleToCutPath(grid [][]int) bool {
	m, n := len(grid), len(grid[0])

	gridCount := make([][]int, m)
	memory := make([][]int, m)
	for i := 0; i < m; i++ {
		gridCount[i] = make([]int, n)
		memory[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memory[i][j] = -1
		}
	}

	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		if i >= m || j >= n {
			return 0
		}

		if i == m-1 && j == n-1 {
			return 1
		}

		if grid[i][j] == 0 {
			return 0
		}

		if memory[i][j] != -1 {
			return memory[i][j]
		}

		r1 := dfs(i+1, j)
		r2 := dfs(i, j+1)
		gridCount[i][j] = r1 + r2

		memory[i][j] = r1 + r2
		return r1 + r2
	}

	count := dfs(0, 0)
	gridCount[0][0] = 1

	if count == 0 {
		return true
	}

	fmt.Println(count)
	fmt.Println(gridCount)

	gridCount2 := make([][]int, m)
	for i := 0; i < m; i++ {
		gridCount2[i] = make([]int, n)
		gridCount2[0][0] = 1
		for j := 0; j < n; j++ {
			if gridCount[i][j] == 0 || grid[i][j] == 0 {
				continue
			}
			if i-1 >= 0 {
				gridCount2[i][j] += gridCount2[i-1][j]
			}
			if j-1 >= 0 {
				gridCount2[i][j] += gridCount2[i][j-1]
			}
			fmt.Println(gridCount2)
			if (i != 0 || j != 0) && gridCount2[i][j]*gridCount[i][j] == count {
				return true
			}
		}
	}

	fmt.Println(gridCount2)
	return false
}
