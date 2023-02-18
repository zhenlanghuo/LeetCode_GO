package main

import "fmt"

func main() {
	fmt.Println(takeCharacters("aabaaaacaabc", 2))
	fmt.Println(takeCharacters("aabaaaacaabc", 2))
}

func takeCharacters(s string, k int) int {
	if k == 0 {
		return 0
	}

	m := make(map[byte]int)
	for _, b := range s {
		m[byte(b)]++
	}

	if len(m) != 3 {
		return -1
	}

	for _, count := range m {
		if count < k {
			return -1
		}
	}

	//fmt.Println(m)

	result := len(s)
	i, j := -1, 0
	for i < len(s)-1 {
		for j < len(s) && m[s[j]] > k {
			m[s[j]]--
			j++
		}

		//fmt.Println(i, j, m)
		result = min(result, i+1+len(s)-j)

		i++
		m[s[i]]++
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
