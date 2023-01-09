package main

import (
	"fmt"
)

func main() {
	fmt.Println(countPalindromicSubsequences("abcdabcdabcdabcdabcdabcdabcdabcddcbadcbadcbadcbadcbadcbadcbadcba"))
}

func countPalindromicSubsequences(s string) int {
	const mod int = 1e9 + 7
	n := len(s)
	dp := make([][][]int, 4)
	for i := 0; i < 4; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, n)
		}
	}

	for i := 0; i < n; i++ {
		dp[s[i]-'a'][i][i] = 1
	}

	for k := 1; k < n; k++ {
		for i := 0; i < n-k; i++ {
			j := i + k
			for c := 0; c < 4; c++ {
				if int(s[i]-'a') == c && int(s[j]-'a') == c {
					dp[c][i][j] = (dp[0][i+1][j-1] + dp[1][i+1][j-1] + dp[2][i+1][j-1] + dp[3][i+1][j-1] + 2) % mod
				} else if int(s[i]-'a') == c {
					dp[c][i][j] = dp[c][i][j-1]
				} else if int(s[j]-'a') == c {
					dp[c][i][j] = dp[c][i+1][j]
				} else {
					dp[c][i][j] = dp[c][i+1][j-1]
				}
			}
		}
	}

	result := 0
	for i := 0; i < 4; i++ {
		result = (result + dp[i][0][n-1]) % mod
	}

	//fmt.Println(dp)

	return result
}

//func countPalindromicSubsequences(s string) int {
//	const mod int = 1e9 + 7
//	n := len(s)
//	dp := make([][]int, n)
//	next := make([][]int, n)
//	prev := make([][]int, n)
//	post := make([]int, 4)
//	for i := 0; i < n; i++ {
//		dp[i] = make([]int, n)
//		next[i] = make([]int, 4)
//		prev[i] = make([]int, 4)
//		dp[i][i] = 1
//	}
//
//	for i := 0; i < 4; i++ {
//		post[i] = n
//	}
//	for i := n - 1; i >= 0; i-- {
//		for j := 0; j < 4; j++ {
//			next[i][j] = post[j]
//		}
//		post[s[i]-'a'] = i
//	}
//
//	for i := 0; i < 4; i++ {
//		post[i] = -1
//	}
//	for i := 0; i < n; i++ {
//		for j := 0; j < 4; j++ {
//			prev[i][j] = post[j]
//		}
//		post[s[i]-'a'] = i
//	}
//
//	for k := 1; k < n; k++ {
//		for i := 0; i < n-k; i++ {
//			j := i + k
//			if s[i] == s[j] {
//				low := next[i][s[i]-'a']
//				high := prev[j][s[i]-'a']
//				if low < high {
//					dp[i][j] = (dp[i+1][j-1]*2%mod - dp[low+1][high-1]%mod + mod) % mod
//				} else if low == high {
//					dp[i][j] = (dp[i+1][j-1]*2%mod + 1) % mod
//				} else {
//					dp[i][j] = (dp[i+1][j-1]*2%mod + 2) % mod
//				}
//			} else {
//				dp[i][j] = (dp[i+1][j]%mod + dp[i][j-1]%mod - dp[i+1][j-1]%mod + mod) % mod
//			}
//			//dp[i][j] %= mod
//		}
//	}
//	//fmt.Println(dp)
//	return dp[0][n-1]
//}
