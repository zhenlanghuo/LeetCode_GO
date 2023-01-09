package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}

func maxProfit(prices []int) int {
	curProfit := 0
	result := 0
	for i := 1; i < len(prices); i++ {
		curProfit += prices[i] - prices[i-1]
		if curProfit < 0 {
			curProfit = 0
		}
		if result < curProfit {
			result = curProfit
		}
	}

	return result
}

//func maxProfit(prices []int) int {
//	rightMaxPrices := make([]int, len(prices))
//	rightMaxPrices[len(prices)-1] = prices[len(prices)-1]
//	for i := len(prices) - 2; i >= 0; i-- {
//		rightMaxPrices[i] = rightMaxPrices[i+1]
//		if prices[i] > rightMaxPrices[i] {
//			rightMaxPrices[i] = prices[i]
//		}
//	}
//
//	max := 0
//	for i := 0; i < len(prices); i++ {
//		if rightMaxPrices[i]-prices[i] > max {
//			max = rightMaxPrices[i] - prices[i]
//		}
//	}
//
//	return max
//}
