package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minCost([]int{1, 2, 1, 2, 1, 3, 3}, 2))
	fmt.Println(minCost([]int{1, 2, 1, 2, 1}, 2))
	fmt.Println(minCost([]int{1, 2, 1, 2, 1}, 5))
}

func minCost(nums []int, k int) int {
	n := len(nums)

	cost := make([][]int, n)
	unique := make([][]int, n)
	for i := 0; i < n; i++ {
		cost[i] = make([]int, n)
		unique[i] = make([]int, n)
		unique[i][i] = 1
	}

	for i := 0; i < n; i++ {
		state := make([]int, n)
		state[nums[i]] = 1
		for j := i + 1; j < n; j++ {
			unique[i][j] = unique[i][j-1]
			switch state[nums[j]] {
			case 0:
				unique[i][j]++
				state[nums[j]]++
			case 1:
				unique[i][j]--
				state[nums[j]]++
			case 2:
			}
			cost[i][j] = j - i + 1 - unique[i][j]
		}
	}

	memory := make([]int, n)

	var dfs func(start int) int
	dfs = func(start int) int {
		if start == len(nums) {
			return 0
		}

		if memory[start] != 0 {
			return memory[start]
		}

		result := math.MaxInt64
		for i := start; i < len(nums); i++ {
			result = min(result, dfs(i+1)+cost[start][i]+k)
		}
		memory[start] = result

		return result
	}

	return dfs(0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
