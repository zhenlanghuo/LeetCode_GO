package main

import "fmt"

func main() {
	fmt.Println(minimumDeletions("aababbab"))
	fmt.Println(minimumDeletions("aabbaaaaabbabb"))
}

func minimumDeletions(s string) int {

	n := len(s)
	aCountLeftToRight := make([]int, n+1)
	aCountRightToLeft := make([]int, n+1)
	for i := 0; i < n; i++ {
		aCountLeftToRight[i+1] = aCountLeftToRight[i]
		if s[i] == 'a' {
			aCountLeftToRight[i+1]++
		}
	}

	for i := n - 1; i >= 0; i-- {
		aCountRightToLeft[i] = aCountRightToLeft[i+1]
		if s[i] == 'a' {
			aCountRightToLeft[i]++
		}
	}

	fmt.Println(aCountLeftToRight)
	fmt.Println(aCountRightToLeft)

	ans := min(aCountLeftToRight[n], n-aCountLeftToRight[n])
	//fmt.Println(ans)

	for i := 0; i < n; i++ {
		//fmt.Println(i, aCountLeftToRight[i+1], aCountRightToLeft[i+1], i+1-aCountLeftToRight[i+1]+aCountRightToLeft[i+1])
		ans = min(ans, i+1-aCountLeftToRight[i+1]+aCountRightToLeft[i+1])
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
