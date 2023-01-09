package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
}

func maxProfit(prices []int) int {

	n := len(prices)
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1]
		if prices[i]-prices[i-1] > 0 {
			dp[i] += prices[i] - prices[i-1]
		}
	}

	return dp[n-1]
}
