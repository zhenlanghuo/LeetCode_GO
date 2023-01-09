package main

import (
	"fmt"
	"math"

	//"sort"
)

func main() {
	fmt.Println(findRotateSteps("edcba", "abcde"))
}

func findRotateSteps(ring string, key string) int {
	n, m := len(ring), len(key)

	// dp[i][j] 代表ring第i个字符对准12点方向开始拼写key第j个(包含j)之后的字符需要的最少操作数
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j < m; j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	ringCharIndexes := make(map[byte][]int)
	for i := 0; i < n; i++ {
		if _, ok := ringCharIndexes[ring[i]]; !ok {
			ringCharIndexes[ring[i]] = make([]int, 0)
		}
		ringCharIndexes[ring[i]] = append(ringCharIndexes[ring[i]], i)
	}

	//for _, index := range ringCharIndexes[key[m-1]] {
	//	dp[index][m-1] = 1
	//}

	for j := m - 1; j >= 0; j-- {
		for i := 0; i < n; i++ {
			targetByte := key[j]
			for _, targetIndex := range ringCharIndexes[targetByte] {
				step := minStep(targetIndex, i, n)
				dp[i][j] = min(dp[i][j], dp[targetIndex][j+1]+step+1)
			}
		}
	}

	return dp[0][0]
}

//func findRotateSteps(ring string, key string) int {
//	n, m := len(ring), len(key)
//
//	ringCharIndexes := make(map[byte][]int)
//	for i := 0; i < n; i++ {
//		if _, ok := ringCharIndexes[ring[i]]; !ok {
//			ringCharIndexes[ring[i]] = make([]int, 0)
//		}
//		ringCharIndexes[ring[i]] = append(ringCharIndexes[ring[i]], i)
//	}
//	fmt.Println(ringCharIndexes)
//
//	// dp[i][j] 代表指针拼写出key前j个字符时ring的第i个字符对准12点方向的最少操作数
//	dp := make([][]int, n)
//	for i := 0; i < n; i++ {
//		dp[i] = make([]int, m)
//		for j := 0; j < m; j++ {
//			dp[i][j] = math.MaxInt32
//		}
//	}
//
//	for _, i := range ringCharIndexes[key[0]] {
//		dp[i][0] = minStep(i, 0, n) + 1
//		//fmt.Println(dp[i][0])
//	}
//
//	//fmt.Println(dp)
//
//	for j := 1; j < m; j++ {
//		targetByte := key[j]
//		prevByte := key[j-1]
//		for _, i := range ringCharIndexes[targetByte] {
//			for _, prevIndex := range ringCharIndexes[prevByte] {
//				step := minStep(i, prevIndex, n)
//				dp[i][j] = min(dp[i][j], dp[prevIndex][j-1]+step+1)
//			}
//		}
//	}
//
//	//fmt.Println(dp)
//
//	result := math.MaxInt32
//	for i := 0; i < n; i++ {
//		result = min(result, dp[i][m-1])
//	}
//
//	return result
//}

func minStep(index, targetIndex, mod int) int {
	return min((index-targetIndex+mod)%mod, (targetIndex-index+mod)%mod)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
