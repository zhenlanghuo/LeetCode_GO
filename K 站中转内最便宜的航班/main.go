package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findCheapestPrice(4, [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}, 0, 3, 1))
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	// dp[i][k]代表从i开始恰好经过k站中转后到dst需要的最少费用
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	//fmt.Println(dp)

	srcMap := make(map[int][][]int)
	for _, flight := range flights {
		src_ := flight[0]
		dst_ := flight[1]
		if _, ok := srcMap[src_]; !ok {
			srcMap[src_] = make([][]int, 0)
		}
		srcMap[src_] = append(srcMap[src_], flight)
		if dst_ == dst {
			dp[src_][0] = flight[2]
			//fmt.Println(src_, flight)
		}
	}

	//fmt.Println(dp)

	for k_ := 1; k_ <= k; k_++ {
		for i := 0; i < n; i++ {
			dp[i][k_] = dp[i][k_-1]
			for _, flight := range srcMap[i] {
				_, dst_, cost := flight[0], flight[1], flight[2]
				dp[i][k_] = min(dp[i][k_], dp[dst_][k_-1]+cost)
			}
		}
	}

	//fmt.Println(dp)

	result := dp[src][k]
	//for k_ := 0; k_ <= k; k_++ {
	//	result = min(result, dp[src][k_])
	//}

	if result == math.MaxInt32 {
		return -1
	}
	return result
}

//func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
//	dpm := make([][]int, n)
//	for i := 0; i < n; i++ {
//		dpm[i] = make([]int, k+1)
//	}
//
//	srcMap := make(map[int][][]int)
//	for _, flight := range flights {
//		src_ := flight[0]
//		if _, ok := srcMap[src_]; !ok {
//			srcMap[src_] = make([][]int, 0)
//		}
//		srcMap[src_] = append(srcMap[src_], flight)
//	}
//
//	result := dp(src, dst, flights, k, dpm, srcMap)
//	if result == math.MaxInt32 {
//		return -1
//	}
//	return result
//}
//
//func dp(src, dst int, flights [][]int, k int, dpm [][]int, srcMap map[int][][]int) int {
//	if k < 0 {
//		return math.MaxInt32
//	}
//
//	if dpm[src][k] != 0 {
//		return dpm[src][k]
//	}
//
//	result := math.MaxInt32
//	for _, flight := range srcMap[src] {
//		if flight[0] != src {
//			continue
//		}
//
//		if flight[1] == dst {
//			result = min(result, flight[2])
//		}
//		result = min(result, flight[2]+dp(flight[1], dst, flights, k-1, dpm, srcMap))
//	}
//
//	dpm[src][k] = result
//	return result
//}
//
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
