package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumRecolors("WBBWWBBWBW", 7))
	fmt.Println(minimumRecolors("BBBWBBWWBBWBW", 7))
	fmt.Println(minimumRecolors("WBWBBBW", 2))
}

func minimumRecolors(blocks string, k int) int {
	n := len(blocks)
	wCount := make([]int, n+1)

	for i, b := range blocks {
		wCount[i+1] = wCount[i]
		if b == 'W' {
			wCount[i+1]++
		}
	}

	ans := math.MaxInt64
	for i := k - 1; i < n; i++ {
		ans = min(ans, wCount[i+1]-wCount[i-k+1])
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
