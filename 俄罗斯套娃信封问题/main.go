package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxEnvelopes([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}))
}

//func maxEnvelopes(envelopes [][]int) int {
//	// dp[i] 表示信封 envelopes[i]最多可以套多少个信封
//	dp := make([]int, len(envelopes))
//	for i := 0; i < len(envelopes); i++ {
//		dp[i] = -1
//	}
//
//	result := 0
//	var dfs func(i int) int
//	dfs = func(i int) int {
//		if dp[i] != -1 {
//			return dp[i]
//		}
//
//		maxn := math.MinInt32
//		for j := 0; j < len(envelopes); j++ {
//			if envelopes[i][0] > envelopes[j][0] && envelopes[i][1] > envelopes[j][1] {
//				maxn = max(maxn, dfs(j))
//			}
//		}
//		dp[i] = max(maxn+1, 1)
//		if result < dp[i] {
//			result = dp[i]
//		}
//		return dp[i]
//	}
//
//	for i := 0; i < len(envelopes); i++ {
//		dfs(i)
//	}
//
//	fmt.Println(dp)
//
//	return result
//}

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || (a[0] == b[0] && a[1] > b[1])
	})

	dp := make([]int, 0)

	for i := 0; i < len(envelopes); i++ {
		h := envelopes[i][1]
		index := sort.SearchInts(dp, h)
		if index >= len(dp) {
			dp = append(dp, h)
		} else {
			dp[index] = h
		}
	}

	return len(dp)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
