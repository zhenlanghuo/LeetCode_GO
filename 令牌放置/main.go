package main

import "sort"

func main() {

}

func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)
	n := len(tokens)
	ans, point := 0, 0
	l, r := 0, n-1
	for l <= r {
		if power > tokens[l] {
			point++
			power -= tokens[l]
			l++
		} else if point > 0 {
			power += tokens[r]
			r--
			point--
		} else {
			break
		}
		ans = max(ans, point)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
