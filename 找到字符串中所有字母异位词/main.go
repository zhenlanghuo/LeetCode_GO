package main

import "fmt"

func main() {
	fmt.Println(findAnagrams("abab", "ab"))
}

func findAnagrams(s string, p string) []int {
	count := make([]int, 26)
	for _, b := range p {
		count[b-'a']--
	}

	left, right := 0, 0
	result := make([]int, 0)
	for right < len(s) {
		count[s[right]-'a']++
		for count[s[right]-'a'] > 0 {
			count[s[left]-'a']--
			left++
		}
		if right-left+1 == len(p) {
			result = append(result, left)
		}
		right++
	}

	return result
}
