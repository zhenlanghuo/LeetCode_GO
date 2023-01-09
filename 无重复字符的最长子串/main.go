package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {

	// 滑动窗口
	sRune := []rune(s)

	temp := make(map[rune]int)

	maxCount := 0
	rIndex := -1
	for i := 0; i < len(sRune); i++ {
		if i != 0 {
			r := sRune[i-1]
			if _, ok := temp[r]; ok {
				delete(temp, r)
			}
		}

		for rIndex+1 < len(sRune) && temp[sRune[rIndex+1]] == 0 {
			temp[sRune[rIndex+1]] = 1
			rIndex = rIndex + 1
		}
		if maxCount < rIndex-i+1 {
			maxCount = rIndex - i + 1
		}
	}

	//sRune := []rune(s)
	//
	//temp := make(map[rune]int)
	//
	//maxCount := 0
	//for i := 0; i < len(sRune); i++ {
	//	r := sRune[i]
	//	pre, ok := temp[r]
	//	if !ok {
	//		temp[r] = i
	//	} else {
	//		temp = make(map[rune]int)
	//		i = pre
	//	}
	//
	//	if len(temp) > maxCount {
	//		maxCount = len(temp)
	//	}
	//}
	//
	return maxCount
}
