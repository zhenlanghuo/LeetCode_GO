package main

import (
	"fmt"
	"sort"
)

func main() {

}

func matrixSum(nums [][]int) int {
	ans := 0

	n, m := len(nums), len(nums[0])
	for i := 0; i < n; i++ {
		sort.Ints(nums[i])
	}

	fmt.Println(nums)

	for j := 0; j < m; j++ {
		temp := 0
		for i := 0; i < n; i++ {
			temp = max(temp, nums[i][j])
		}
		ans += temp
	}

	//for i := n - 1; i >= 0; i-- {
	//	temp := 0
	//	for j := 0; j < m; j++ {
	//		temp = max(temp, nums[i][j])
	//	}
	//	ans += temp
	//}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
