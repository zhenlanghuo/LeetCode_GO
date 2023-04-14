package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minScoreTriangulation([]int{1, 3, 1, 4, 1, 5}))
}

func minScoreTriangulation(values []int) int {

	n := len(values)

	memory := make([][]int, n)
	for i := 0; i < n; i++ {
		memory[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memory[i][j] = -1
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == j-1 {
			return 0
		}
		res := math.MaxInt64

		if memory[i][j] != -1 {
			return memory[i][j]
		}

		for k := i + 1; k < j; k++ {
			res = min(res, dfs(i, k)+dfs(k, j)+values[i]*values[j]*values[k])
		}
		memory[i][j] = res

		return res
	}

	return dfs(0, n-1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
