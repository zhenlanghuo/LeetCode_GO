package main

import "sort"

func main() {

}

func buyChoco(prices []int, money int) int {
	sort.Ints(prices)

	ans := money

	sum := 0
	for i := 0; i < len(prices); i++ {
		if i > 1 {
			break
		}
		sum += prices[i]
	}
	if ans >= sum {
		ans -= sum
	}

	return ans
}
