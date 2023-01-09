package main

import "fmt"

func main() {
	fmt.Println(numRollsToTarget(2, 6, 7))
}

func numRollsToTarget(n int, k int, target int) int {
	const mod = 1e9 + 7
	//// dp[i][j]表示前i个骰子正面朝上的数字和为j的组合数
	//dp := make([][]int, n+1)
	//for i := 0; i <= n; i++ {
	//	dp[i] = make([]int, target+1)
	//}
	//dp[0][0] = 1
	//
	//for i := 1; i <= n; i++ {
	//	for h := 1; h <= target; h++ {
	//		for j := 1; j <= h && j <= k; j++ {
	//			dp[i][h] = (dp[i][h] + dp[i-1][h-j]) % mod
	//		}
	//	}
	//}
	//
	//fmt.Println(dp)
	//return dp[n][target]

	// 空间压缩
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for h := target; h >= 0; h-- {
			// 需要先置0, 因为这是每组必须选一个的分组背包,
			dp[h] = 0
			for j := 1; j <= h && j <= k; j++ {
				dp[h] = (dp[h] + dp[h-j]) % mod
			}
		}
	}

	//fmt.Println(dp)
	return dp[target]
}
