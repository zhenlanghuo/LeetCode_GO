package main

import (
	"fmt"
)

func main() {
	fmt.Println(lastSubstring("leetcode"))
	fmt.Println(lastSubstring("abab"))
}

func lastSubstring(s string) string {

	n := len(s)
	i, j, k := 0, 1, 0
	for i+k < n && j+k < n {
		if s[i+k] == s[j+k] {
			k++
		} else if s[i+k] < s[j+k] {
			i = i + k + 1
			if i >= j {
				j = i + 1
			}
			k = 0
		} else {
			j = j + k + 1
			k = 0
		}
	}

	return s[i:]
}
