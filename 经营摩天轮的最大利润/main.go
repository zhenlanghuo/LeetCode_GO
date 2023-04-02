package main

import "fmt"

func main() {
	fmt.Println(minOperationsMaxProfit([]int{8, 3}, 5, 6))
	fmt.Println(minOperationsMaxProfit([]int{10, 9, 6}, 6, 4))
	fmt.Println(minOperationsMaxProfit([]int{3, 4, 0, 5, 1}, 1, 92))
}

func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	profit := 0
	maxProfit := 0
	ans := -1
	wait := 0
	i := 0
	for {
		if i < len(customers) {
			wait += customers[i]
		}

		if wait > 0 {
			profit += min(wait, 4)*boardingCost - runningCost
			wait -= min(wait, 4)
		}

		//fmt.Println(i, profit, wait)

		if profit > maxProfit {
			maxProfit = profit
			ans = i + 1
		}
		i++

		if wait == 0 && i >= len(customers) {
			break
		}

	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
