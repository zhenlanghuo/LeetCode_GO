package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{5, 4, 3, 2, 1}))
}

func maxProfit(prices []int) int {
	k := 2
	n := len(prices)
	buy := make([][]int, n)
	sell := make([][]int, n)
	for i := 0; i < n; i++ {
		buy[i] = make([]int, k+1)
		sell[i] = make([]int, k+1)
		for j := 1; j <= k; j++ {
			buy[i][j] = math.MinInt32
			sell[i][j] = math.MinInt32
		}
	}
	buy[0][1] = -prices[0]

	for i := 1; i < n; i++ {
		//sell[i][0] =
		for j := 1; j <= k; j++ {
			buy[i][j] = max(buy[i-1][j], sell[i-1][j-1]-prices[i])
			sell[i][j] = max(sell[i-1][j], buy[i-1][j]+prices[i])
		}
		fmt.Println(buy)
		fmt.Println(sell)
	}

	//fmt.Println(buy)
	//fmt.Println(sell)

	result := 0
	for i := 0; i <= k; i++ {
		result = max(result, sell[n-1][i])
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
