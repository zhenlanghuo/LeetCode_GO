package main

import "fmt"

func main() {
	fmt.Println(waysToReachTarget(6, [][]int{{6, 1}, {3, 2}, {2, 3}}))
}

func waysToReachTarget(target int, types [][]int) int {
	n := len(types)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, target+1)
		dp[i][0] = 1
	}

	for i := 1; i <= n; i++ {
		count, marks := types[i-1][0], types[i-1][1]
		//fmt.Println(i, count, marks)
		//for j := 1; j <= target; j++ {
		//	dp[i][j] = dp[i-1][j]
		//}
		for j := 1; j <= target; j++ {
			dp[i][j] = dp[i-1][j]
			for k := 1; k <= count; k++ {
				if j >= k*marks {
					dp[i][j] = (dp[i][j] + dp[i-1][j-k*marks]) % (1e9 + 7)
				}
			}
		}
	}

	fmt.Println(dp)

	return dp[n][target]
}
