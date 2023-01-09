package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(superEggDrop(100, 8191))
}

//func superEggDrop(k int, n int) int {
//	// dp[i][j] 代表有j个鸡蛋可操作i次,最多可以探测的楼层数
//	dp := make([][]int, n+1)
//	for i := 0; i <= n; i++ {
//		dp[i] = make([]int, k+1)
//		dp[i][1] = i
//	}
//
//	for i := 1; dp[i-1][k] < n; i++ {
//		for j := 2; j <= k; j++ {
//			// 如果鸡蛋没有碎，那么这一层的上面可以有dp[i-1][j]层
//			// 如果鸡蛋碎了，那么这一层的下面可以有dp[i-1][j-1]层
//			dp[i][j] = 1 + dp[i-1][j-1] + dp[i-1][j]
//		}
//	}
//
//	result := math.MaxInt32
//	for i := 0; i <= n; i++ {
//		if dp[i][k] >= n {
//			result = min(result, i)
//		}
//	}
//
//	return result
//}

func superEggDrop(k int, n int) int {

	memory := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		memory[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			memory[i][j] = -1
		}
	}

	// dp[i][j] 代表有j个鸡蛋可操作i次,最多可以探测的楼层数
	var dp func(k int, n int) int

	dp = func(k int, n int) int {
		if memory[k][n] != -1 {
			return memory[k][n]
		}

		if n == 0 {
			memory[k][n] = 0
			return 0
		}

		if k == 1 {
			memory[k][n] = n
			return n
		}

		//lo := 1
		//hi := n

		memory[k][n] = 1 + dp(k-1, n-1) + dp(k, n-1)
		return memory[k][n]
	}

	dp(k, n)

	result := math.MaxInt32
	for i := 0; i <= n; i++ {
		if memory[k][i] >= n {
			result = min(result, i)
		}
	}

	fmt.Println(memory)

	return result
}

//func superEggDrop(k int, n int) int {
//
//	memory := make([][]int, k+1)
//	for i := 0; i <= k; i++ {
//		memory[i] = make([]int, n+1)
//		for j := 0; j <= n; j++ {
//			memory[i][j] = -1
//		}
//	}
//
//	// k个鸡蛋n层最少的操作次数
//	var dp func(k int, n int) int
//
//	dp = func(k int, n int) int {
//		if memory[k][n] != -1 {
//			return memory[k][n]
//		}
//
//		if n == 0 {
//			return 0
//		}
//		if k == 1 {
//			return n
//		}
//
//		lo := 1
//		hi := n
//		result := math.MaxInt32
//		for lo <= hi {
//			mid := (lo + hi) / 2
//			broken := dp(k-1, mid-1) + 1
//			noBroken := dp(k, n-mid) + 1
//			if broken == noBroken {
//				result = min(result, broken)
//				break
//			} else if broken > noBroken {
//				result = min(result, broken)
//				hi = mid - 1
//			} else {
//				result = min(result, noBroken)
//				lo = mid + 1
//			}
//		}
//
//		//for i := n; i >= 1; i-- {
//		//	result = min(result, max(dp(k, n-i), dp(k-1, i-1))+1)
//		//}
//		memory[k][n] = result
//
//		return result
//	}
//
//	return dp(k, n)
//}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
