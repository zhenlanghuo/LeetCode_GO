package main

import "fmt"

func main() {
	fmt.Println(findLongestWord("aaa", []string{"aaa", "aa", "a"}))
}

func findLongestWord(s string, dictionary []string) string {

	dp := make([][26]int, len(s)+1)
	for i := 0; i < 26; i++ {
		dp[len(s)][i] = len(s)
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			if s[i] == byte(j+'a') {
				dp[i][j] = i
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}


	result := ""
	for _, str := range dictionary {
		if !isSubStr(dp, str) {
			continue
		}

		if len(str) > len(result) {
			result = str
		}

		if len(result) == len(str) && result > str {
			result = str
		}
	}

	return result
}

func isSubStr(dp [][26]int, str string) bool {
	index := 0
	for _, b := range str {
		index = dp[index][b-'a']
		if index == len(dp)-1 {
			return false
		}
		index++
	}

	return true
}
