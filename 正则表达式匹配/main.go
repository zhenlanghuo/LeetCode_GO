package main

import "fmt"

func main() {
	fmt.Println(isMatch("aab", "c*a*b"))
}

func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	memory := make([][]int, n)
	for i := 0; i < n; i++ {
		memory[i] = make([]int, m)
		for j := 0; j < m; j++ {
			memory[i][j] = -1
		}
	}

	// s[0:i+1] 和 p[0:j+1] 是否匹配
	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		if i == -1 && j == -1 {
			return true
		}

		if i == -1 {
			// 检查是否剩余n对的*号字符匹配
			if j > 0 && j%2 == 1 {
				for j_ := 1; j_ <= j; j_+=2 {
					if p[j_] != '*' {
						return false
					}
				}
				return true
			}
		}

		if j < 0 || i < 0 {
			return false
		}

		if memory[i][j] != -1 {
			return memory[i][j] == 1
		}

		result := false
		if s[i] == p[j] || p[j] == '.' {
			result = result || dp(i-1, j-1)
		} else if p[j] == '*' {
			// 匹配0个
			result = result || dp(i, j-2)
			// 匹配多个
			if s[i] == p[j-1] || p[j-1] == '.' {
				result = result || dp(i-1, j)
			}
		}

		memory[i][j] = 0
		if result {
			memory[i][j] = 1
		}

		return result
	}

	return dp(n-1, m-1)
}

//func isMatch(s string, p string) bool {
//	n, m := len(s), len(p)
//	dp := make([][]bool, n+1)
//	for i := 0; i <= n; i++ {
//		dp[i] = make([]bool, m+1)
//	}
//	dp[0][0] = true
//
//	for i := 0; i <= n; i++ {
//		for j := 1; j <= m; j++ {
//			if i >= 1 && (p[j-1] == '.' || s[i-1] == p[j-1]) {
//				dp[i][j] = dp[i][j] || dp[i-1][j-1]
//			} else if p[j-1] == '*' {
//				dp[i][j] = dp[i][j] || dp[i][j-2]
//				if i >= 1 && (s[i-1] == p[j-2] || p[j-2] == '.') {
//					dp[i][j] = dp[i][j] || dp[i-1][j]
//				}
//			}
//		}
//	}
//
//	fmt.Println(dp)
//
//	return dp[n][m]
//}

//
//func isMatch(s string, p string) bool {
//	dp := make([][]bool, len(s)+1)
//	for i := 0; i < len(s)+1; i++ {
//		dp[i] = make([]bool, len(p)+1)
//	}
//
//	dp[0][0] = true
//	for i := 0; i < len(s)+1; i++ {
//		for j := 1; j < len(p)+1; j++ {
//			if p[j-1] == byte('*') {
//				if i >= 1 && (s[i-1] == p[j-2] || p[j-2] == byte('.')) {
//					dp[i][j] = dp[i-1][j] || dp[i][j-2]
//				} else {
//					dp[i][j] = dp[i][j-2]
//				}
//			} else if i >= 1 && (p[j-1] == s[i-1] || p[j-1] == byte('.')) {
//				dp[i][j] = dp[i-1][j-1]
//			}
//		}
//	}
//
//	//fmt.Println(dp)
//
//	return dp[len(s)][len(p)]
//}
