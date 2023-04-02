package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(answerQueries([]int{4, 5, 2, 1}, []int{3, 10, 21}))
}

func answerQueries(nums []int, queries []int) []int {

	sort.Ints(nums)
	n := len(nums)
	sums := make([]int, n+1)
	for i, num := range nums {
		sums[i+1] = sums[i] + num
	}

	m := len(queries)
	ans := make([]int, m)
	for i, query := range queries {
		firstLTIndex := sort.SearchInts(sums, query)
		ans[i] = firstLTIndex
		if firstLTIndex > n || sums[firstLTIndex] > query {
			ans[i] = firstLTIndex - 1
		}
	}

	return ans
}
