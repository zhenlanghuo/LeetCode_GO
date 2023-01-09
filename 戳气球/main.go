package main

import "fmt"

func main() {
	fmt.Println(maxCoins([]int{3, 1, 5, 8}))
}

func maxCoins(nums []int) int {
	val := make([]int, len(nums)+2)
	val[0], val[len(nums)+1] = 1, 1
	for i := 1; i <= len(nums); i++ {
		val[i] = nums[i-1]
	}

	fmt.Println(val)

	rec := make([][]int, len(nums)+2)
	for i := 0; i < len(rec); i++ {
		rec[i] = make([]int, len(nums)+2)
		for j := 0; j < len(rec[i]); j++ {
			rec[i][j] = -1
		}
	}
	//dfs(val, 0, len(nums)+1, rec)
	//fmt.Println(rec)

	for k := 2; k <= len(nums)+1; k++ {
		for left := 0; left+k < len(nums)+2; left++ {
			right := left + k
			for i := left + 1; i < right; i++ {
				sum := val[left]*val[right]*val[i] + max(rec[left][i], 0) + max(rec[i][right], 0)
				rec[left][right] = max(rec[left][right], sum)
			}
		}
	}

	//fmt.Println(rec)

	return rec[0][len(nums)+1]
}

func dfs(val []int, left, right int, rec [][]int) int {
	if left+1 >= right {
		return 0
	}

	if rec[left][right] != -1 {
		return rec[left][right]
	}

	for i := left + 1; i < right; i++ {
		sum := val[left]*val[right]*val[i] + dfs(val, left, i, rec) + dfs(val, i, right, rec)
		rec[left][right] = max(rec[left][right], sum)
	}

	return rec[left][right]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
