package main

import "fmt"

func main() {
	fmt.Println(canPartition([]int{14, 9, 8, 4, 3, 2}))
}

//func canPartition(nums []int) bool {
//	n := len(nums)
//
//	sum := 0
//	for i := 0; i < n; i++ {
//		sum += nums[i]
//	}
//
//	if sum%2 != 0 {
//		return false
//	}
//
//	dp := make([][]bool, n+1)
//	for i := 0; i <= n; i++ {
//		dp[i] = make([]bool, sum/2+1)
//	}
//	dp[0][0] = true
//
//	for i := 1; i <= n; i++ {
//		for j := 0; j <= sum/2; j++ {
//			dp[i][j] = dp[i-1][j]
//			if j-nums[i-1] >= 0 {
//				if dp[i-1][j-nums[i-1]] {
//					dp[i][j] = true
//				}
//			}
//			if dp[i][sum/2] {
//				return true
//			}
//		}
//	}
//
//	fmt.Println(dp)
//
//	return false
//}

func canPartition(nums []int) bool {
	n := len(nums)

	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false
	}

	dp := make([]bool, sum/2+1)
	dp[0] = true
	for i := 0; i < n; i++ {
		for j := sum / 2; j >= 1; j-- {
			if j-nums[i] >= 0 {
				if dp[j-nums[i]] {
					dp[j] = true
				}
				if dp[sum/2] {
					return true
				}
			}
		}
	}
	fmt.Println(dp)

	return false
}
