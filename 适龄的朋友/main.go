package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(numFriendRequests([]int{20, 30, 100, 110, 120}))
	//fmt.Println(numFriendRequests([]int{16, 16, 16, 16}))
	fmt.Println(numFriendRequests([]int{108, 115, 5, 24, 82}))
}

func numFriendRequests(ages []int) int {
	n := len(ages)
	sort.Ints(ages)
	reverse(ages)
	fmt.Println(ages)
	r, ans := 0, 0
	lastL := 0
	for l, age := range ages {
		for r < n && ages[r] <= age && age < 2*ages[r]-14 {
			r++
		}
		//fmt.Println(l, r)
		ans += max(r-l-1, 0)
		if l > 0 && ages[l-1] != age {
			lastL = l
		}
		if l > 0 && ages[l-1] == age && age < 2*age-14 {
			ans += l - lastL
		}
	}

	return ans
}

func reverse(array []int) {
	l, r := 0, len(array)-1
	for l < r {
		array[l], array[r] = array[r], array[l]
		l++
		r--
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
