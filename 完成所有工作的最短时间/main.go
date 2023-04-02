package main

import "math"

func main() {

}

func minimumTimeRequired(jobs []int, k int) int {
	n := len(jobs)
	sum := make([]int, 1<<uint(n))
	dp := make([][]int, k)
	for i := 0; i < k; i++ {
		dp[i] = make([]int, 1<<uint(n))
	}

	for i := 0; i < n; i++ {
		for j := 1; j < 1<<uint(n); j++ {
			if 1<<uint(i)&j > 0 {
				sum[j] += jobs[i]
			}
		}
	}

	//fmt.Println(sum)
	dp[0] = sum
	for i := 1; i < k; i++ {
		for j := 1; j < 1<<uint(n); j++ {
			dp[i][j] = math.MaxInt64
			for s := j; s > 0; s = (s - 1) & j {
				dp[i][j] = min(dp[i][j], max(dp[i-1][j^s], sum[s]))
			}
		}
	}

	//fmt.Println(dp)

	return dp[k-1][1<<uint(n)-1]
}

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
