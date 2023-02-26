package main

import "fmt"

func main() {
	fmt.Println(leftRigthDifference([]int{10, 4, 8, 3}))
}

func leftRigthDifference(nums []int) []int {
	n := len(nums)
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	//fmt.Println(sum)

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		//fmt.Println(sum[i], sum[n]-sum[i+1])
		ans[i] = abs(sum[i] - (sum[n] - sum[i+1]))
	}

	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
