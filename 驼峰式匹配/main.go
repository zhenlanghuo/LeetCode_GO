package main

import "fmt"

func main() {
	fmt.Println(camelMatch([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FB"))
	fmt.Println(camelMatch([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FoBa"))
	fmt.Println(camelMatch([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FoBaT"))
}

func camelMatch(queries []string, pattern string) []bool {
	n := len(queries)
	ans := make([]bool, n)
	for i := 0; i < n; i++ {
		ans[i] = true
	}
	for i, query := range queries {
		qIndex, pIndex := 0, 0
		for qIndex < len(query) {
			if pIndex < len(pattern) && query[qIndex] == pattern[pIndex] {
				qIndex++
				pIndex++
				continue
			}

			if query[qIndex]-'A' <= 26 {
				ans[i] = false
				break
			}
			qIndex++
		}
		if pIndex != len(pattern) {
			ans[i] = false
		}
	}

	return ans
}
