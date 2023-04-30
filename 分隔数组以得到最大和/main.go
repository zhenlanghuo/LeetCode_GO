package main

import "fmt"

func main() {
	fmt.Println(maxSumAfterPartitioning([]int{1, 15, 7, 9, 2, 5, 10}, 3))
	fmt.Println(maxSumAfterPartitioning([]int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}, 4))
}

func maxSumAfterPartitioning(arr []int, k int) int {

	n := len(arr)
	memory := make([]int, n)
	for i := 0; i < n; i++ {
		memory[i] = -1
	}

	// arr[start:]分成k个子数组的元素最大和
	var dfs func(start int) int
	dfs = func(start int) int {
		if start >= n {
			return 0
		}
		if memory[start] != -1 {
			return memory[start]
		}
		ans := 0
		mx := 0
		for i := start; i < n && i < start+k; i++ {
			mx = max(mx, arr[i])
			ans = max(ans, mx*(i-start+1)+dfs(i+1))
		}
		memory[start] = ans
		return ans
	}

	ans := dfs(0)
	fmt.Println(memory)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
