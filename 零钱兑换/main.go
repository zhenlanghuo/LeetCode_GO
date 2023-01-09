package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}

func coinChange(coins []int, amount int) int {

	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
		for _, coin := range coins {
			if i < coin {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}

//func coinChange(coins []int, amount int) int {
//	if amount == 0 {
//		return 0
//	}
//
//	dp := make([]int, amount+1)
//	for i := 0; i <= amount; i++ {
//		dp[i] = amount + 1
//	}
//
//	dp[0] = 0
//	for i := 1; i <= amount; i++ {
//		for j := 0; j < len(coins); j++ {
//			if i < coins[j] {
//				continue
//			}
//			dp[i] = min(dp[i], dp[i-coins[j]]+1)
//		}
//	}
//
//	if dp[amount] == amount+1 {
//		return -1
//	}
//
//	return dp[amount]
//}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
