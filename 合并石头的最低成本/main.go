package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mergeStones([]int{3, 2, 4, 1}, 2))
	fmt.Println(mergeStones([]int{3, 2, 4, 1}, 3))
	fmt.Println(mergeStones([]int{3, 5, 1, 2, 6}, 3))
	fmt.Println(mergeStones([]int{6, 4, 4, 6}, 2))
}

func mergeStones(stones []int, k int) int {
	n := len(stones)

	if (n-1)%(k-1) != 0 {
		return -1
	}

	sum := make([]int, n+1)
	for i, stone := range stones {
		sum[i+1] = sum[i] + stone
	}

	memory := make([][][]int, n)
	for i := 0; i < n; i++ {
		memory[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			memory[i][j] = make([]int, k+1)
			for z := 0; z <= k; z++ {
				memory[i][j][z] = -1
			}
		}
	}
	var dfs func(i, j, p int) int
	dfs = func(i, j, p int) int {
		if memory[i][j][p] != -1 {
			return memory[i][j][p]
		}

		if p == 1 {
			if i == j {
				return 0
			}
			memory[i][j][p] = dfs(i, j, k) + sum[j+1] - sum[i]
			return memory[i][j][p]
		}
		res := math.MaxInt64
		for m := i; m < j; m += k - 1 {
			res = min(res, dfs(i, m, 1)+dfs(m+1, j, p-1))
		}
		memory[i][j][p] = res
		return res
	}
	return dfs(0, n-1, 1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
