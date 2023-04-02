package main

import (
	"fmt"
)

func main() {
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
	fmt.Println(shipWithinDays([]int{3, 2, 2, 4, 1, 4}, 3))
	fmt.Println(shipWithinDays([]int{1, 2, 3, 1, 1}, 4))
}

func shipWithinDays(weights []int, days int) int {

	check := func(maxWeights int) bool {
		temp := 0
		count := 0
		for _, weight := range weights {
			if weight > maxWeights {
				return false
			}
			if temp+weight > maxWeights {
				count++
				temp = weight
			} else {
				temp += weight
			}
		}
		count++
		return count <= days
	}

	sum := 0
	for _, weight := range weights {
		sum += weight
	}

	l, r := 0, sum
	for l <= r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
}

//func shipWithinDays(weights []int, days int) int {
//
//	n := len(weights)
//	// dp[i][j] 代表能在 i+1 天内将传送带上前 j+1 个包裹送达的船的最低运载能力
//	dp := make([][]int, days)
//	for i := 0; i < days; i++ {
//		dp[i] = make([]int, n)
//	}
//
//	sum := make([]int, n+1)
//	for i := 0; i < n; i++ {
//		sum[i+1] = sum[i] + weights[i]
//	}
//
//	dp[0] = sum[1:]
//	for i := 1; i < days; i++ {
//		for j := 0; j < n; j++ {
//			dp[i][j] = math.MaxInt64
//			for k := 0; k <= j; k++ {
//				dp[i][j] = min(dp[i][j], max(dp[i-1][k], sum[j+1]-sum[k+1]))
//			}
//		}
//	}
//
//	return dp[days-1][n-1]
//}
//
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
