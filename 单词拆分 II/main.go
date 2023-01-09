package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
}

var result []string

func wordBreak(s string, wordDict []string) []string {
	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}
	dp := make([]bool, len(s)+1)
	record := make(map[int][]int, 0)
	dp[0] = true
	for i := 1; i < len(s)+1; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
				record[i] = append(record[i], j)
			}
		}
	}

	fmt.Println(dp)
	fmt.Println(record)

	result = make([]string, 0)
	buildString(record, s, nil, len(s))

	return result
}

func buildString(record map[int][]int, s string, tempList []string, index int) {
	if index == 0 {
		clone := make([]string, len(tempList))
		copy(clone, tempList)
		reverse(clone)
		result = append(result, strings.Join(clone, " "))
		fmt.Println(strings.Join(clone, " "))
		//return
	}

	prevList := record[index]
	for _, prevIndex := range prevList {
		buildString(record, s, append(tempList, s[prevIndex:index]), prevIndex)
	}
}

func reverse(tempList []string) {
	start := 0
	end := len(tempList) - 1
	for start < end {
		tempList[start], tempList[end] = tempList[end], tempList[start]
		start++
		end--
	}
}
