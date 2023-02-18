package main

import "fmt"

func main() {
	//fmt.Println(countGoodStrings(2, 3, 1, 2))
	//fmt.Println(countGoodStrings(3, 3, 1, 1))
	fmt.Println(countGoodStrings(200, 200, 10, 1))
	fmt.Println(countGoodStrings(20, 20, 1, 6))

	//fmt.Println(Factorial(1,1))
}

func countGoodStrings(low int, high int, zero int, one int) int {
	ans := 0
	dp := make([]int, high+1)
	dp[0] = 1
	for i := 1; i <= high; i++ {
		if i-zero >= 0 {
			dp[i] = (dp[i] + dp[i-zero]) % (1e9 + 7)
		}

		if i-one >= 0 {
			dp[i] = (dp[i] + dp[i-one]) % (1e9 + 7)
		}

		if i >= low {
			ans = (ans + dp[i]) % (1e9 + 7)
		}
	}

	return ans
}
