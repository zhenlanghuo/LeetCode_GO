package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(bestTeamScore([]int{1, 3, 5, 10, 15}, []int{1, 2, 3, 4, 5}))
	fmt.Println(bestTeamScore([]int{4, 5, 6, 5}, []int{2, 1, 2, 1}))
	fmt.Println(bestTeamScore([]int{1, 2, 3, 5}, []int{8, 9, 10, 1}))
}

func bestTeamScore(scores []int, ages []int) int {
	type Pair struct {
		score int
		age   int
	}

	n := len(scores)
	pairs := make([]*Pair, n)
	for i := 0; i < n; i++ {
		pairs[i] = &Pair{
			score: scores[i],
			age:   ages[i],
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].score == pairs[j].score {
			return pairs[i].age < pairs[j].age
		}
		return pairs[i].score < pairs[j].score
	})

	ans := 0
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = pairs[i].score
		for j := i - 1; j >= 0; j-- {
			if pairs[i].age >= pairs[j].age {
				dp[i] = max(dp[i], dp[j]+pairs[i].score)
			}
		}
		ans = max(ans, dp[i])
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
