package main

import "fmt"

func main() {
	//fmt.Println(tallestBillboard([]int{1, 2, 3, 6}))
	fmt.Println(tallestBillboard([]int{1, 2, 3, 4, 5, 6}))
	//fmt.Println(tallestBillboard([]int{1, 2}))
}

func tallestBillboard(rods []int) int {
	n := len(rods)
	// 对于每一条钢筋, 有三种选择:不选择、乘+1、乘-1
	// key为被选择的钢筋的总和, val为最大的正数和
	dp := make([]map[int]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make(map[int]int)
	}
	dp[0] = map[int]int{0: 0}

	result := 0
	for i, rod := range rods {
		for k, v := range dp[i] {
			dp[i+1][k+rod] = max(dp[i+1][k+rod], v+rod)
			dp[i+1][k-rod] = max(dp[i+1][k-rod], v)
			dp[i+1][k] = max(dp[i+1][k], v)
		}
		result = max(result, dp[i+1][0])
	}

	//fmt.Println(dp)

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
