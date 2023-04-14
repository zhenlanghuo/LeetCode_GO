package main

import "fmt"

func main() {
	//fmt.Println(longestDecomposition("ghiabcdefhelloadamhelloabcdefghi"))
	//fmt.Println(longestDecomposition("merchant"))
	fmt.Println(longestDecomposition("antaprezatepzapreanta"))
}

func longestDecomposition(text string) int {
	if len(text) == 0 {
		return 0
	}

	n := len(text)
	for i := 1; i <= n/2; i++ {
		if text[:i] == text[n-i:] {
			return 2 + longestDecomposition(text[i:n-i])
		}
	}
	return 1
}
