package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(miceAndCheese([]int{1, 1, 3, 4}, []int{4, 4, 1, 1}, 2))
	fmt.Println(miceAndCheese([]int{1, 1}, []int{1, 1}, 2))
}

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	n := len(reward1)
	type pair struct {
		diff  int
		index int
	}
	diffs := make([]*pair, n)
	for i := 0; i < n; i++ {
		diffs[i] = &pair{diff: reward1[i] - reward2[i], index: i}
	}

	sort.Slice(diffs, func(i, j int) bool {
		return diffs[i].diff > diffs[j].diff
	})

	ans := 0
	for i := 0; i < k; i++ {
		ans += reward1[diffs[i].index]
		reward2[diffs[i].index] = 0
	}

	for _, num := range reward2 {
		ans += num
	}

	return ans
}
