package main

import "fmt"

func main() {
	fmt.Println(canCross([]int{0, 1, 3, 5, 6, 8, 12, 17}))
}

func canCross(stones []int) bool {
	n := len(stones)
	dp := make([][]bool, n)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n)
		if i >= 1 && stones[i]-stones[i-1] > i {
			return false
		}
	}
	dp[0][0] = true
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			k := stones[i] - stones[j]
			if k > j+1 {
				break
			}
			dp[i][k] = dp[j][k-1] || dp[j][k] || dp[j][k+1]
			if i == n-1 && dp[i][k] {
				return true
			}
		}
	}

	//fmt.Println(dp)

	return false
}

//func canCross(stones []int) bool {
//	if stones[1] != 1 {
//		return false
//	}
//
//	m := len(stones)
//	dp := make(map[int]map[int]bool)
//
//	//for i := 0; i < len(dp); i++ {
//	//	dp[i] = make([]bool, stones[m-1]+2)
//	//}
//
//	nums := []int{-1, 0, 1}
//
//	for _, num := range nums {
//		_, ok := dp[1+num+1]
//		if !ok {
//			dp[1+num+1] = make(map[int]bool)
//		}
//		dp[1+num+1][num+1] = true
//	}
//
//	//fmt.Println(dp)
//
//	for i := 2; i < m; i++ {
//		if _, ok := dp[stones[i]]; !ok {
//			continue
//		}
//		for k, ok := range dp[stones[i]] {
//			if !ok {
//				continue
//			}
//			for _, num := range nums {
//				_, ok = dp[stones[i]+k+num]
//				if !ok {
//					dp[stones[i]+k+num] = make(map[int]bool)
//				}
//				dp[stones[i]+k+num][k+num] = true
//			}
//		}
//	}
//
//	//fmt.Println(dp)
//
//	if _, ok := dp[stones[m-1]]; ok {
//		return true
//	}
//
//	return false
//}
