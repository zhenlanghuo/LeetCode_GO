package main

import "fmt"

func main() {
	fmt.Println(divisibilityArray("998244353", 3))
	fmt.Println(divisibilityArray("91221181269244172125025075166510211202115152121212341281327", 21))
}

func divisibilityArray(word string, m int) []int {
	n := len(word)
	remainder := 0
	ans := make([]int, 0, n)
	for _, b := range word {
		remainder = (remainder*10 + int(b-'0')) % m
		if remainder == 0 {
			ans = append(ans, 1)
		} else {
			ans = append(ans, 0)
		}
	}

	return ans
}
