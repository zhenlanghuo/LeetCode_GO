package main

import "fmt"

func main() {
	fmt.Println(change(10, []int{1, 2, 5}))
}

//func change(amount int, coins []int) int {
//	dp := make([][]int, len(coins)+1)
//	for i := 0; i < len(coins)+1; i++ {
//		dp[i] = make([]int, amount+1)
//		dp[i][0] = 1
//	}
//
//	for i := 1; i < len(coins)+1; i++ {
//		coin := coins[i-1]
//		for j := 1; j <= amount; j++ {
//			dp[i][j] += dp[i-1][j]
//			for k := coin; k <= j; k += coin {
//				dp[i][j] += dp[i-1][j-k]
//			}
//		}
//	}
//
//	fmt.Println(dp)
//
//	return dp[len(coins)][amount]
//}

//func change(amount int, coins []int) int {
//	dp := make([][]int, len(coins)+1)
//	for i := 0; i < len(coins)+1; i++ {
//		dp[i] = make([]int, amount+1)
//		dp[i][0] = 1
//	}
//
//	for i := 1; i < len(coins)+1; i++ {
//		coin := coins[i-1]
//		for j := 1; j <= amount; j++ {
//			dp[i][j] += dp[i-1][j]
//			if j >= coin {
//				dp[i][j] += dp[i][j-coin]
//			}
//		}
//	}
//
//	// fmt.Println(dp)
//
//	return dp[len(coins)][amount]
//}

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	fmt.Println(dp)

	return dp[amount]
}
