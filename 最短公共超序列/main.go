package main

import (
	"fmt"
)

func main() {
	fmt.Println(shortestCommonSupersequence("abac", "cab"))
	fmt.Println(shortestCommonSupersequence("aaaaaaaa", "aaaaaaaa"))
}

type Pair struct {
	length int
	i      int
	j      int
}

func shortestCommonSupersequence(str1 string, str2 string) string {

	n, m := len(str1), len(str2)
	dp := make([][]Pair, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]Pair, m+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j].length = dp[i-1][j-1].length + 1
				dp[i][j].i = i - 1
				dp[i][j].j = j - 1
			} else {
				if dp[i-1][j].length > dp[i][j-1].length {
					dp[i][j].length = dp[i-1][j].length
					dp[i][j].i = i - 1
					dp[i][j].j = j
				} else {
					dp[i][j].length = dp[i][j-1].length
					dp[i][j].i = i
					dp[i][j].j = j - 1
				}
			}
		}
	}

	fmt.Println(dp)
	lcs := make([]byte, 0, dp[n][m].length)

	temp := dp[n][m]
	i, j := n-1, m-1
	for {
		if i < n && j < m && str1[i] == str2[j] {
			lcs = append(lcs, str1[i])
		}
		i, j = temp.i-1, temp.j-1
		temp = dp[temp.i][temp.j]
		if len(lcs) == dp[n][m].length {
			break
		}
	}

	reverse(lcs)
	fmt.Println(string(lcs))

	result := make([]byte, 0, n+m-dp[n][m].length)
	i, j = 0, 0
	k := 0
	for i < n && j < m && k < len(lcs) {
		if str1[i] != str2[j] {
			if str1[i] != lcs[k] {
				result = append(result, str1[i])
				i++
			}
			if str2[j] != lcs[k] {
				result = append(result, str2[j])
				j++
			}
		} else {
			result = append(result, str1[i])
			if str1[i] == lcs[k] {
				k++
			}
			i++
			j++
		}
	}

	for i < n {
		result = append(result, str1[i])
		i++
	}

	for j < m {
		result = append(result, str2[j])
		j++
	}

	return string(result)
}

func reverse(a []byte) {
	l, r := 0, len(a)-1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
}
