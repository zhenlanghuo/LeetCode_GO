package main

import "fmt"

func main() {
	fmt.Println(minDistance("horse", "ros"))
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}

	for i := 1; i <= n1; i++ {
		dp[i][0] = i
	}

	for i := 1; i <= n2; i++ {
		dp[0][i] = i
	}

	// dp[i-1][j-1]+1 代表进行替换操作
	// dp[i-1][j]+1 代表进行删除操作
	// dp[i][j-1]+1 代表进行插入操作
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j]+1, dp[i][j-1]+1))
			} else {
				dp[i][j] = min(dp[i-1][j-1]+1, min(dp[i-1][j]+1, dp[i][j-1]+1))
			}
		}
	}

	return dp[n1][n2]
}

//func minDistance(word1 string, word2 string) int {
//
//	dp := make([][]int, 0, len(word1)+1)
//	for i := 0; i <= len(word1); i++ {
//		dp = append(dp, make([]int, len(word2)+1))
//	}
//
//	for i := 0; i <= len(word1); i++ {
//		dp[i][0] = i
//	}
//
//	for j := 0; j <= len(word2); j++ {
//		dp[0][j] = j
//	}
//
//	for i := 1; i <= len(word1); i++ {
//		for j := 1; j <= len(word2); j++ {
//			if word1[i-1] == word2[j-1] {
//				dp[i][j] = min(dp[i][j-1]+1, dp[i-1][j]+1)
//				dp[i][j] = min(dp[i][j], dp[i-1][j-1])
//			} else {
//				dp[i][j] = min(dp[i][j-1]+1, dp[i-1][j]+1)
//				dp[i][j] = min(dp[i][j], dp[i-1][j-1]+1)
//			}
//		}
//	}
//
//	return dp[len(word1)][len(word2)]
//}
