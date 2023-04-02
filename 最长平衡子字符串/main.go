package main

import "fmt"

func main() {
	fmt.Println(findTheLongestBalancedSubstring("01000111"))
	fmt.Println(findTheLongestBalancedSubstring("00111"))
	fmt.Println(findTheLongestBalancedSubstring("111"))
}

func findTheLongestBalancedSubstring(s string) int {
	ans := 0
	i := 0
	n := len(s)
	for i < n {
		zeroCount := 0
		for i < n && s[i] == '0' {
			i++
			zeroCount++
		}
		//fmt.Println(zeroCount, i)
		oneCount := 0
		for i < n && s[i] == '1' {
			i++
			oneCount++
		}
		//fmt.Println(oneCount, i)
		ans = max(ans, min(zeroCount, oneCount)*2)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
