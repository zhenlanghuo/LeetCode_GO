package main

import "fmt"

func main() {
	//fmt.Println(maximumOr([]int{12, 9}, 1))
	//fmt.Println(maximumOr([]int{1, 8, 2}, 2))
	fmt.Println(maximumOr([]int{10, 8, 4}, 1))
}

func maximumOr(nums []int, k int) int64 {
	n := len(nums)

	dp := make([][]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int64, k+1)
	}

	dp[0][0] = int64(nums[0])
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] | int64(nums[i])
	}
	for i := 1; i <= k; i++ {
		dp[0][i] = dp[0][i-1] * 2
	}

	fmt.Println(dp)

	for i := 1; i < n; i++ {
		for j := 1; j < k+1; j++ {
			//dp[i][j] = max(max(dp[i-1][j]|int64(nums[i]), dp[i-1][j-1]|int64(nums[i]*2)), dp[i][j-1]|int64(nums[i]*2))
			//dp[i][j] = max(dp[i-1][j]|int64(nums[i]), dp[i-1][j-1]|int64(nums[i]*2))
			dp[i][j] = dp[i-1][j] | int64(nums[i])
			temp := nums[i]
			for h := 1; h <= j; h++ {
				temp = temp * 2
				dp[i][j] = max(dp[i][j], dp[i-1][j-h]|int64(temp))
			}
		}
	}

	fmt.Println(dp)

	return dp[n-1][k]
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
