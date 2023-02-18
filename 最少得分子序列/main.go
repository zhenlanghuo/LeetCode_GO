package main

import "fmt"

func main() {
	//fmt.Println(minimumScore("abacaba", "bzaa"))
	//fmt.Println(minimumScore("cde", "xyz"))
	//fmt.Println(minimumScore("adebddaccdcabaade", "adbae"))
	fmt.Println(minimumScore("gbjbacdiiiecgceeafdcdhjhhcjfchjbejibhejgjhhhjhifahfbbcfcajehjgcjgffjdejbhjahgffgdaifhhgaadjiabfdf", "hidefgbgjghceagdaaib"))
}

func minimumScore(s string, t string) int {
	pre := make([]int, len(s)+1)
	suf := make([]int, len(s)+1)

	i, j := 0, 0
	for i < len(s) {
		pre[i+1] = pre[i]
		if j < len(t) && s[i] == t[j] {
			j++
			pre[i+1]++
		}
		i++
	}
	//fmt.Println(pre)

	i, j = len(s)-1, len(t)-1
	for i >= 0 {
		suf[i] = suf[i+1]
		if j >= 0 && s[i] == t[j] {
			j--
			suf[i]++
		}
		i--
	}
	//fmt.Println(suf)

	// i = -1时, 表示s全部拿来匹配后缀
	// i = len(s)-1时, 表示s全部拿来匹配前缀
	i = -1
	result := len(t)
	for ; i < len(s); i++ {
		temp := len(t) - (pre[i+1] + suf[i+1])
		if temp < 0 {
			temp = 0
		}
		if temp >= 0 && result > temp {
			result = temp
		}
	}

	return result
}
