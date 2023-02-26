package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minImpossibleOR([]int{2, 1}))
}

func minImpossibleOR(nums []int) int {
	sort.Ints(nums)
	need := 1
	for _, num := range nums {
		if num == need {
			need = need * 2
		}
		if num > need {
			break
		}
	}

	return need
}
