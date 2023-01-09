package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 3, 4, 5, 6, 7, 8}, 3))
}

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	buy := make([]int, n)
	sell := make([]int, n)
	buy[0] = -prices[0]

	for i := 1; i < n; i++ {
		buy[i] = max(buy[i-1], sell[i-1]-prices[i])
		sell[i] = max(sell[i-1], buy[i-1]+prices[i]-fee)
	}

	//fmt.Println(buy)
	//fmt.Println(sell)

	return sell[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
