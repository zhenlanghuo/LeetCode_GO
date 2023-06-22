package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(paintWalls([]int{1, 2, 3, 2}, []int{1, 2, 3, 2}))
	fmt.Println(paintWalls([]int{42, 8, 28, 35, 21, 13, 21, 35}, []int{2, 1, 1, 1, 2, 1, 1, 2}))
	//fmt.Println(paintWalls([]int{2, 3, 4, 2}, []int{1, 1, 1, 1}))
}

func paintWalls(cost []int, time []int) int {

	ans := math.MaxInt32
	n := len(cost)
	allTime := 0
	for i := 0; i < n; i++ {
		allTime += time[i]
	}
	allTime++

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, allTime)
		for j := 0; j < allTime; j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	//for j := 0; j < allTime; j++ {
	//	dp[0][j] = 0
	//}
	for i := 0; i <= n; i++ {
		dp[i][0] = 0
	}
	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		for j := time[i-1]; j < allTime; j++ {
			//dp[i][j] = dp[i-1][j]
			dp[i][j] = min(dp[i-1][j], dp[i-1][j-time[i-1]]+cost[i-1])
		}
	}

	if n%2 != 0 {
		n++
	}

	fmt.Println(dp)
	fmt.Println(n)
	for j := n / 2; j < allTime; j++ {
		ans = min(ans, dp[n][j])
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
