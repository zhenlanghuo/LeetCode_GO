package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(makeArrayIncreasing([]int{1, 5, 3, 6, 7}, []int{1, 3, 2, 4}))
	fmt.Println(makeArrayIncreasing([]int{1, 5, 3, 6, 7}, []int{4, 3, 1}))
	fmt.Println(makeArrayIncreasing([]int{1, 5, 3, 6, 7}, []int{1, 6, 3, 3}))
	fmt.Println(makeArrayIncreasing([]int{1, 2, 3, 6, 7}, []int{1, 6, 3, 3}))
}

func makeArrayIncreasing(arr1 []int, arr2 []int) int {

	n := len(arr1)
	// dp[i][j] 表示 arr1[:i]经过j次替换之后，arr1[:i]为严格递增时arr1[i-1]的最小值, 如果arr1[:i]没法经过j次替换变为严格递增, dp[i][j]=-1
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			dp[i][j] = math.MaxInt64
		}
	}
	dp[1][0] = arr1[0]
	dp[0][0] = -1

	//fmt.Println(dp)

	sort.Ints(arr2)
	fmt.Println(arr2)

	for i := 1; i < n+1; i++ {
		for j := 0; j <= i && j < n+1; j++ {
			if j-1 >= 0 {
				mn := dp[i-1][j-1]

				if mn != math.MaxInt64 {
					// 换
					index := sort.SearchInts(arr2, mn+1)
					if index < len(arr2) {
						dp[i][j] = arr2[index]
					}
				}
			}

			// 不换
			if dp[i-1][j] != math.MaxInt64 && arr1[i-1] > dp[i-1][j] && arr1[i-1] < dp[i][j] {
				dp[i][j] = arr1[i-1]
			}
		}
	}

	//fmt.Println(dp)

	for i := 0; i < n+1; i++ {
		if dp[n][i] != math.MaxInt64 {
			return i
		}
	}

	return -1
}
