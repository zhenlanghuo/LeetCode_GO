package main

import "fmt"

func main() {
	fmt.Println(PredictTheWinner([]int{1, 5, 233, 7}))
}

func PredictTheWinner(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true
	}
	// dp[i][j] 代表先手能拿到nums[i:j+1]最大的值
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}

	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] += sum[i-1] + nums[i-1]
	}

	fmt.Println(sum)

	for k := 1; k < n; k++ {
		for i := 0; i < n-k; i++ {
			j := i + k
			sumij := sum[j+1] - sum[i]
			// dp[i][j] 最大值取决于 dp[i+1][j] dp[i][j-1] 的大小
			// 当处于[i,j+1]的情况时,先手拿nums[i],那么能拿到最大值的取决于处于[i+1,j]的情况先手能拿的最大值
			dp[i][j] = max(sumij-dp[i+1][j], sumij-dp[i][j-1])
		}
	}

	fmt.Println(dp[0][n-1])

	if dp[0][n-1]*2 >= sum[n] {
		return true
	}

	return false
}

//func PredictTheWinner(nums []int) bool {
//	if len(nums) == 1 {
//		return true
//	}
//
//	sum := 0
//	for _, num := range nums {
//		sum += num
//	}
//
//	n := len(nums)
//	memory := make([]int, n+1)
//
//	var dp func(nums []int, n int, sum int) int
//	dp = func(nums []int, n int, sum int) int {
//		if n == 1 {
//			return nums[0]
//		}
//
//		result1 := sum - dp(nums[1:n], n-1, sum-nums[0])
//		result2 := sum - dp(nums[0:n-1], n-1, sum-nums[n-1])
//
//		memory[n] = max(result1, result2)
//		return memory[n]
//	}
//
//	dp(nums, n, sum)
//
//	//fmt.Println(memory)
//
//	if memory[n]*2 >= sum {
//		return true
//	}
//
//	return false
//}
//
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
