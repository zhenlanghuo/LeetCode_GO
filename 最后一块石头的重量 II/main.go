package main

import "fmt"

func main() {
	fmt.Println(lastStoneWeightII([]int{31, 26, 33, 21, 40}))
}

//func lastStoneWeightII(stones []int) int {
//	sum := 0
//	n := len(stones)
//	for i := 0; i < n; i++ {
//		sum += stones[i]
//	}
//
//	// 寻找石头子集的重量和不超过target的最大重量和
//	target := sum / 2
//
//	//// dp[i][j] 表示前i个石头是否能找出重量和为j的子集
//	//dp := make([][]bool, n+1)
//	//for i := 0; i <= n; i++ {
//	//	dp[i] = make([]bool, target+1)
//	//	dp[i][0] = true
//	//}
//
//	//for i := 1; i <= n; i++ {
//	//	for j := 1; j <= target; j++ {
//	//		dp[i][j] = dp[i-1][j]
//	//		if j-stones[i-1] >= 0 {
//	//			dp[i][j] = dp[i][j] || dp[i-1][j-stones[i-1]]
//	//		}
//	//	}
//	//}
//	//
//	//maxTarget := 0
//	//for i := target; i >= 0; i-- {
//	//	if dp[n][i] {
//	//		maxTarget = i
//	//		break
//	//	}
//	//}
//
//	// dp[i][j] 表示前i个石头是否能找出重量和为j的子集
//	// 空间压缩为1维
//	dp := make([]bool, target+1)
//	dp[0] = true
//
//	for i := 1; i <= n; i++ {
//		for j := target; j >= stones[i-1]; j-- {
//			dp[j] = dp[j] || dp[j-stones[i-1]]
//		}
//	}
//
//	maxTarget := 0
//	for i := target; i >= 0; i-- {
//		if dp[i] {
//			maxTarget = i
//			break
//		}
//	}
//
//	return sum - 2*maxTarget
//}

func lastStoneWeightII(stones []int) int {
	sum := 0
	n := len(stones)
	for i := 0; i < n; i++ {
		sum += stones[i]
	}

	// 寻找石头子集的重量和不超过target的最大重量和
	target := sum / 2

	//// dp[i][j] 表示前i个石头中重量和不超过j的最大重量和
	//dp := make([][]int, n+1)
	//for i := 0; i <= n; i++ {
	//	dp[i] = make([]int, target+1)
	//	//dp[i][0] = 0
	//}
	//
	//for i := 1; i <= n; i++ {
	//	for j := 1; j <= target; j++ {
	//		dp[i][j] = dp[i-1][j]
	//		if j-stones[i-1] >= 0 {
	//			dp[i][j] = max(dp[i][j], dp[i-1][j-stones[i-1]]+stones[i-1])
	//		}
	//	}
	//}

	//return sum - 2*dp[n][target]

	// 空间压缩为1维
	dp := make([]int, target+1)
	for i := 1; i <= n; i++ {
		for j := target; j >= stones[i-1]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i-1]]+stones[i-1])
		}
	}

	return sum - 2*dp[target]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
