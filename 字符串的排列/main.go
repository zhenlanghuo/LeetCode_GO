package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("ab", "a"))
}

func checkInclusion(s1 string, s2 string) bool {
	cnt := make([]int, 26)
	for _, b := range s1 {
		cnt[b-'a']--
	}

	left := 0
	for right, b := range s2 {
		cnt[b-'a']++
		for cnt[b-'a'] > 0 {
			cnt[s2[left]-'a']--
			left++
		}
		if right-left+1 == len(s1) {
			return true
		}
	}

	return false
}
