package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestStrChain([]string{"a", "b", "ba", "bca", "bda", "bdca"}))
	fmt.Println(longestStrChain([]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}))

}

func longestStrChain(words []string) int {

	m := make(map[string]int)
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	//fmt.Println(words)
	ans := 0
	for _, word := range words {
		m[word] = 1
		for i := 0; i < len(word); i++ {
			preWord := word[0:i] + word[i+1:]
			m[word] = max(m[word], 1+m[preWord])
			ans = max(ans, m[word])
			//fmt.Println(word, preWord, m)
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
