package main

import "fmt"

func main() {
	fmt.Println(wordBreak("catsanddog", []string{"cat","cats","and","sand","dog"}))
}

func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i < len(s)+1; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
			}
		}
	}

	return dp[len(s)]
}
