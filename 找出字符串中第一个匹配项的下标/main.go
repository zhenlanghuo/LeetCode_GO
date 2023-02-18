package main

import "fmt"

func main() {
	//"aabaaabaaac"
	//"aabaaac"
	fmt.Println(strStr("aaaabbaaaaaacc", "ababcdababab"))
}

//func strStr(haystack string, needle string) int {
//	pi := make([]int, len(needle))
//	for i := 1; i < len(needle); i++ {
//		j := i
//		for j > 0 {
//			j = pi[j-1]
//			if needle[j] == needle[i] {
//				pi[i] = j + 1
//				break
//			}
//		}
//	}
//	fmt.Println(pi)
//	return 0
//}

func strStr(haystack string, needle string) int {
	pi := make([]int, len(needle))
	pi[0] = 0

	// next函数求解的解释
	// 如何更好地理解和掌握 KMP 算法? - 阮行止的回答 - 知乎
	// https://www.zhihu.com/question/21923021/answer/1032665486
	for i, j := 1, 0; i < len(needle); i++ {
		for j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}

	fmt.Println(pi)

	i, j := 0, 0
	for i < len(haystack) {
		for i < len(haystack) && haystack[i] == needle[j] {
			i++
			j++
			if j == len(needle) {
				return i - j
			}
		}

		if i >= len(haystack) {
			break
		}

		if j == 0 {
			i++
		} else {
			//i += j - pi[j-1] + pi[j-1]
			j = pi[j-1]
		}

	}

	return -1
}
