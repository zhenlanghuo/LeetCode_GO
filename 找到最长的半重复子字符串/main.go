package main

import "fmt"

func main() {
	//fmt.Println(longestSemiRepetitiveSubstring("52233"))
	//fmt.Println(longestSemiRepetitiveSubstring("5494"))
	//fmt.Println(longestSemiRepetitiveSubstring("52234665"))
	fmt.Println(longestSemiRepetitiveSubstring("1111111"))
}

func longestSemiRepetitiveSubstring(s string) int {
	n := len(s)

	if n <= 2 {
		return n
	}

	mx := 2
	l, r := 0, 1
	count := 0
	if s[l] == s[r] {
		count = 1
	}

	//fmt.Println(n)

	for r < n-1 {
		//fmt.Println(l, r, count)
		if s[r] != s[r+1] {
			r++
			mx = max(mx, r-l+1)
		} else {
			if count == 1 {
				for {
					l++
					if s[l-1] == s[l] {
						break
					}
				}
				count--
			} else {
				r++
				mx = max(mx, r-l+1)
				count++
			}
		}
	}
	mx = max(mx, r-l+1)

	return mx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
