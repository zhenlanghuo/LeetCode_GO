package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(minExtraChar("leetscode", []string{"leet", "code", "leetcode"}))
	//fmt.Println(minExtraChar("leetscodeab", []string{"leetscode", "leet", "codeab"}))
	fmt.Println(minExtraChar("sayhelloworld", []string{"hello", "world"}))
}

func minExtraChar(s string, dictionary []string) int {
	n := len(s)
	sbytes := []byte(s)

	dictionaryBytes := make([][]byte, len(dictionary))
	for i := 0; i < len(dictionary); i++ {
		dictionaryBytes[i] = []byte(dictionary[i])
	}

	m := make([]int, n)
	for i := 0; i < len(m); i++ {
		m[i] = -1
	}

	var dfs func(start int) int
	dfs = func(start int) int {
		if start >= n {
			return 0
		}

		if m[start] != -1 {
			return m[start]
		}

		mn := math.MaxInt64
		for _, dBytes := range dictionaryBytes {
			if start+len(dBytes) > n {
				continue
			}

			if compare(sbytes[start:start+len(dBytes)], dBytes) {
				mn = min(mn, dfs(start+len(dBytes)))
			}
		}

		mn = min(mn, 1+dfs(start+1))
		mn = min(mn, n-start)
		m[start] = mn
		return mn
	}

	return dfs(0)
}

func compare(b1, b2 []byte) bool {
	//fmt.Println(string(b1), string(b2))
	for i := 0; i < len(b1); i++ {
		if b1[i] != b2[i] {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
