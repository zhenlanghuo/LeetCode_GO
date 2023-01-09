package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{1}))
}

func maxProfit(prices []int) int {
	n := len(prices)
	// buy[i] 代表第i天，手上有股票的最大收益
	buy := make([]int, n)
	// sell[i][0] 代表第i天，手上没有股票且处于冷冻期的最大收益
	// sell[i][1] 代表第i天，手上没有股票且不处于冷冻期的最大收益
	sell := make([][2]int, n)

	buy[0] = -prices[0]
	// 第0天没有进行过交易, 不可能处于冷冻期, 属于非法状态，置为一个非常小的值
	sell[0][0] = math.MinInt32
	sell[0][1] = 0

	for i := 1; i < n; i++ {
		buy[i] = max(buy[i-1], sell[i-1][1]-prices[i])
		sell[i][0] = buy[i-1] + prices[i]
		sell[i][1] = max(sell[i-1][0], sell[i-1][1])
	}

	//fmt.Println(buy)

	//fmt.Println(sell)

	return max(sell[n-1][0], sell[n-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
