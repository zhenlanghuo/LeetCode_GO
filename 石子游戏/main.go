package main

import "fmt"

func main() {
	fmt.Println(stoneGame([]int{3, 7, 2, 3}))
}

func stoneGame(piles []int) bool {

	n := len(piles)
	if n == 1 {
		return true
	}
	// dp[i][j] 代表先手能拿到nums[i:j+1]最大的值
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = piles[i]
	}

	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] += sum[i-1] + piles[i-1]
	}

	//fmt.Println(sum)

	for k := 1; k < n; k++ {
		for i := 0; i < n-k; i++ {
			j := i + k
			sumij := sum[j+1] - sum[i]
			// dp[i][j] 最大值取决于 dp[i+1][j] dp[i][j-1] 的大小
			// 当处于[i,j+1]的情况时,先手拿nums[i],那么能拿到最大值的取决于处于[i+1,j]的情况先手能拿的最大值
			dp[i][j] = max(sumij-dp[i+1][j], sumij-dp[i][j-1])
		}
	}

	//fmt.Println(dp[0][n-1])

	if dp[0][n-1]*2 >= sum[n] {
		return true
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
