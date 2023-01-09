package main

import "fmt"

func main() {
	fmt.Println(profitableSchemes(5, 3, []int{2, 2}, []int{2, 3}))
	fmt.Println(profitableSchemes(10, 5, []int{2, 3, 5}, []int{6, 7, 8}))
}

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	const mod = 1e9 + 7
	sz := len(group)

	//// dp[i][j][k] 表示前i个可选择的计划里最多可用j人的情况下利润至少为k得组合计划数
	//dp := make([][][]int, sz+1)
	//for i := 0; i <= sz; i++ {
	//	dp[i] = make([][]int, n+1)
	//	for j := 0; j <= n; j++ {
	//		dp[i][j] = make([]int, minProfit+1)
	//		dp[0][j][0] = 1
	//	}
	//}
	//
	//for i := 1; i <= sz; i++ {
	//	for j := 0; j <= n; j++ {
	//		for k := 0; k <= minProfit; k++ {
	//			dp[i][j][k] = dp[i-1][j][k]
	//			if j-group[i-1] >= 0 {
	//				// 当需要的利润为负数时, 归到利润为0那去
	//				dp[i][j][k] = (dp[i][j][k] + dp[i-1][j-group[i-1]][max(k-profit[i-1], 0)]) % mod
	//			}
	//
	//		}
	//	}
	//}
	//
	////fmt.Println(dp)
	//
	////fmt.Println(dp[1])
	//
	//return dp[sz][n][minProfit]

	// 空间压缩
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, minProfit+1)
		dp[i][0] = 1
	}

	for i := 1; i <= sz; i++ {
		for j := n; j >= group[i-1]; j-- {
			for k := 0; k <= minProfit; k++ {
				dp[j][k] = (dp[j][k] + dp[j-group[i-1]][max(k-profit[i-1], 0)]) % mod
			}
		}
	}

	//fmt.Println(dp)

	//fmt.Println(dp[1])

	return dp[n][minProfit]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
