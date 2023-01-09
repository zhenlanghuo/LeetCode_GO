package main

import "fmt"

func main() {
	s := "babgbag"
	t := "bag"
	fmt.Println(numDistinct(s, t))
}

func numDistinct(s string, t string) int {

	dp := make([][]int, len(t), len(t))
	for i := 0; i < len(t); i++ {
		dp[i] = make([]int, len(s), len(s))
	}

	// dp[i][j] 表示 s[:j+1] 的子序列中出现 t[:t+1] 的个数

	// 初始化 dp[0][0] 的数据
	if s[0] == t[0] {
		dp[0][0] = 1
	}
	// 初始化 dp[0][j] 的数据
	for j:=1;j<len(s);j++ {
		if t[0] == s[j] {
			dp[0][j] = dp[0][j-1] + 1
		} else {
			dp[0][j] = dp[0][j-1]
		}
	}

	for i := 1; i < len(t); i++ {
		for j := i; j < len(s); j++ {
			if t[i] == s[j] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[len(t)-1][len(s)-1]
}
