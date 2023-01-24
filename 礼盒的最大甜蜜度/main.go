package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumTastiness([]int{13, 5, 1, 8, 21, 2}, 3))
}
func maximumTastiness(price []int, k int) int {

	sort.Ints(price)
	//fmt.Println(price)

	check := func(d int) bool {
		cnt, x := 1, price[0]
		for _, v := range price {
			if v >= d+x {
				cnt++
				x = v
			}
		}
		return cnt >= k
	}

	left := 0
	right := ((price[len(price)-1] - price[0]) / (k - 1)) + 1
	//fmt.Println(right)

	for left <= right {
		mid := (left + right) / 2
		if check(mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left - 1
}
