package main

import (
	"fmt"
)

func main() {
	fmt.Println(largestNumber([]int{4, 3, 2, 5, 6, 7, 2, 5, 5}, 9))
	//fmt.Println(largestNumber([]int{6, 10, 15, 40, 40, 40, 40, 40, 40}, 47))
}

func largestNumber(cost []int, target int) string {
	n := len(cost)
	// dp[i][j] 代表从前i个数位中选取出总成本为j的最大的数位数量
	dp := make([][]int, n+1)
	from := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, target+1)
		from[i] = make([]int, target+1)
		for j := 0; j <= target; j++ {
			dp[i][j] = -1
		}
	}
	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		for j := 0; j <= target; j++ {
			dp[i][j] = dp[i-1][j]
			from[i][j] = j
			if j >= cost[i-1] && dp[i][j-cost[i-1]] >= 0 {
				dp[i][j] = max(dp[i][j], dp[i][j-cost[i-1]]+1)
				if dp[i][j] == dp[i][j-cost[i-1]]+1 {
					from[i][j] = j - cost[i-1]
				}
			}
		}
	}

	fmt.Println(dp)
	fmt.Println(from)

	if dp[n][target] < 0 {
		return "0"
	}

	bytes := make([]byte, 0, dp[n][target])
	i, j := n, target
	for i > 0 {
		if from[i][j] == j {
			i--
		} else {
			bytes = append(bytes, byte(i+'0'))
			j = from[i][j]
		}
	}

	return string(bytes)

	//// 空间压缩
	//dp := make([]int, target+1)
	//for i := 0; i <= target; i++ {
	//	dp[i] = -1
	//}
	//dp[0] = 0
	//
	//for i := 0; i < n; i++ {
	//	for j := cost[i]; j <= target; j++ {
	//		if dp[j-cost[i]] >= 0 {
	//			dp[j] = max(dp[j], dp[j-cost[i]]+1)
	//		}
	//	}
	//}
	//
	//if dp[target] < 0 {
	//	return "0"
	//}
	//
	//fmt.Println(dp)
	//
	//bytes := make([]byte, 0, dp[target])
	//
	//for i, j := n-1, target; i >= 0 && j >= 0; i-- {
	//	for j >= cost[i] && dp[j] == dp[j-cost[i]]+1 {
	//		bytes = append(bytes, byte(i+1+'0'))
	//		j = j - cost[i]
	//	}
	//}
	//
	//return string(bytes)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
