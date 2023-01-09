package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isAlienSorted([]string{"hello","leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
	fmt.Println(isAlienSorted([]string{"word","world","row"}, "worldabcefghijkmnpqstuvxyz"))
}

func isAlienSorted(words []string, order string) bool {
	orderMap := make(map[byte]int)
	for i, b := range order {
		orderMap[byte(b)] = i
	}

	clone := make([]string, len(words))
	copy(clone, words)

	sort.Slice(clone, func(i, j int) bool {
		minL := min(len(words[i]), len(words[j]))
		for k := 0; k < minL; k++ {
			if words[i][k] != words[j][k] {
				return orderMap[words[i][k]] < orderMap[words[j][k]]
			}
		}
		return len(words[i]) < len(words[j])
	})

	for i := 0; i < len(words); i++ {
		if words[i] != clone[i] {
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
