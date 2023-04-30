package main

import (
	"math"
)

func main() {
	//fmt.Println(minHeightShelves([][]int{{}}))
}

func minHeightShelves(books [][]int, shelfWidth int) int {

	n := len(books)
	memory := make([]int, n)
	for i := 0; i < n; i++ {
		memory[i] = -1
	}

	var dfs func(start int) int
	dfs = func(start int) int {
		if start >= n {
			return 0
		}
		if memory[start] != -1 {
			return memory[start]
		}
		ans := math.MaxInt64
		width := 0
		height := 0
		for i := start; i < n; i++ {
			width += books[i][0]
			if width > shelfWidth {
				break
			}
			height = max(height, books[i][1])
			ans = min(ans, dfs(i+1)+height)
		}
		memory[start] = ans
		return ans
	}

	return dfs(0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
