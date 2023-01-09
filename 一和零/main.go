package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findMaxForm([]string{"10", "0", "1"}, 1, 1))
}

func findMaxForm(strs []string, m int, n int) int {
	l := len(strs)
	//list01 := make([][]int, l+1)
	//for i := 0; i < l; i++ {
	//	list01[i] = make([]int, 2)
	//	for _, c := range strs[i] {
	//		switch c {
	//		case '0':
	//			list01[i][0]++
	//		case '1':
	//			list01[i][1]++
	//		}
	//	}
	//}

	//dp := make([][][]int, l+1)
	//for i := 0; i <= l; i++ {
	//	dp[i] = make([][]int, m+1)
	//	for j := 0; j <= m; j++ {
	//		dp[i][j] = make([]int, n+1)
	//	}
	//}
	//
	////fmt.Println(list01)
	//
	//for i := 1; i <= l; i++ {
	//	num0 := list01[i-1][0]
	//	num1 := list01[i-1][1]
	//	for j := 0; j <= m; j++ {
	//		for k := 0; k <= n; k++ {
	//			dp[i][j][k] = dp[i-1][j][k]
	//			if j-num0 >= 0 && k-num1 >= 0 {
	//				dp[i][j][k] = max(dp[i][j][k], dp[i-1][j-num0][k-num1]+1)
	//			}
	//		}
	//	}
	//}
	//
	////fmt.Println(dp)
	//
	//return dp[l][m][n]

	// 空间压缩
	dp := make([][]int, m+1)
	for j := 0; j <= m; j++ {
		dp[j] = make([]int, n+1)
	}

	//fmt.Println(list01)

	for i := 1; i <= l; i++ {
		num0 := strings.Count(strs[i-1], "0")
		num1 := len(strs[i-1]) - num0
		for j := m; j >= num0; j-- {
			for k := n; k >= num1; k-- {
				dp[j][k] = max(dp[j][k], dp[j-num0][k-num1]+1)
			}
		}
	}

	//fmt.Println(dp)

	return dp[m][n]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
