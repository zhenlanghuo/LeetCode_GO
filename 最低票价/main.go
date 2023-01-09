package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(mincostTickets([]int{1, 2, 3, 4, 6, 8, 9, 10, 13, 14, 16, 17, 19, 21, 24, 26, 27, 28, 29}, []int{3, 14, 50}))
	fmt.Println(mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
}

//func mincostTickets(days []int, costs []int) int {
//	// dp[i]表示days[0, i-1]的最低票价
//	dp := make([]int, len(days)+1)
//	dp[0] = 0
//	//dp[1] = costs[0]
//
//	daysForCosts := []int{1, 7, 30}
//
//	for i := 1; i < len(days)+1; i++ {
//		minn := math.MaxInt32
//		for j := 0; j < len(daysForCosts); j++ {
//			// 找到第一个大于等于days[i-1]-daysForCosts[j]+1的位置x, 那么区间[x,i-1]小于等于daysForCosts[j]
//			// 位置x-1是位置x开始买票cost[j]的位置, 而这里的位置是days数组的位置, 对应的dp数组里的位置要+1, 所以这里prevIndex不需要再-1了
//			prevIndex := sort.SearchInts(days, days[i-1]-daysForCosts[j]+1)
//			if prevIndex < 0 {
//				prevIndex = 0
//			}
//			minn = min(minn, dp[prevIndex]+costs[j])
//		}
//		dp[i] = minn
//	}
//
//	fmt.Println(dp)
//
//	return dp[len(days)]
//}

//func mincostTickets(days []int, costs []int) int {
//	// dp[i]表示days[i:]的最低票价
//	dp := make([]int, len(days))
//	for i := 0; i < len(days); i++ {
//		dp[i] = math.MaxInt32
//
//	}
//	daysForCosts := []int{1, 7, 30}
//
//	var dfs func(start int) int
//
//	dfs = func(start int) int {
//		if start >= len(days) {
//			return 0
//		}
//
//		if dp[start] != math.MaxInt32 {
//			return dp[start]
//		}
//
//		minn := math.MaxInt32
//		for i := 0; i < len(daysForCosts); i++ {
//			nextIndex := sort.SearchInts(days[start:], days[start]+daysForCosts[i])
//			minn = min(minn, dfs(nextIndex+start)+costs[i])
//		}
//		dp[start] = minn
//		return dp[start]
//	}
//
//	//fmt.Println(dp)
//
//	return dfs(0)
//}

func mincostTickets(days []int, costs []int) int {
	// dp[i]表示days[i:]的最低票价
	dp := make([]int, len(days)+1)
	daysForCosts := []int{1, 7, 30}

	for i := len(days) - 1; i >= 0; i-- {
		minn := math.MaxInt32
		for j := 0; j < len(daysForCosts); j++ {
			nextIndex := sort.SearchInts(days[i:], days[i]+daysForCosts[j])
			minn = min(minn, dp[nextIndex+i]+costs[j])
		}
		dp[i] = minn
	}

	fmt.Println(dp)

	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
