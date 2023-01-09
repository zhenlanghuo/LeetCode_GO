package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(hIndex([]int{1,2,100}))
}

func hIndex(citations []int) int {
	max := 0
	for i := 0; i < len(citations); i++ {
		index := sort.SearchInts(citations, i+1)
		if len(citations) - index >= i+1 {
			if max < i+1 {
				max = i+1
			}
		}
	}

	return max
}
