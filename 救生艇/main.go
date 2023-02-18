package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numRescueBoats([]int{1, 2}, 3))
	fmt.Println(numRescueBoats([]int{3, 2, 2, 1}, 3))
	fmt.Println(numRescueBoats([]int{3, 5, 3, 4}, 5))
}

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	ans := 0
	n := len(people)
	l, r := 0, n-1
	for l < r {
		if people[r]+people[l] <= limit {
			l++
		}
		r--
		ans++
	}
	if l == r {
		ans++
	}

	return ans
}
