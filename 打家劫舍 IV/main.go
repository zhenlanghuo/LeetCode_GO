package main

import (
	"fmt"
)

func main() {
	fmt.Println(minCapability([]int{2, 3, 5, 9}, 2))
	fmt.Println(minCapability([]int{2, 7, 9, 3, 1}, 2))
}

func minCapability(nums []int, k int) int {
	check := func(x int) bool {
		f0, f1 := 0, 0
		for _, num := range nums {
			if num <= x {
				f0, f1 = f1, max(f0+1, f1)
			} else {
				f0 = f1
			}
		}

		return f1 >= k
	}

	left := 1
	right := int(1e9)
	for left <= right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

//func minCapability(nums []int, k int) int {
//	n := len(nums)
//
//	memory := make([][]int, n)
//	for i := 0; i < n; i++ {
//		memory[i] = make([]int, k+1)
//		for j := 0; j <= k; j++ {
//			memory[i][j] = -1
//		}
//	}
//
//	var dfs func(start, k int) int
//	dfs = func(start, k int) int {
//		//if start >= n {
//		//	return math.MaxInt64
//		//}
//
//		if memory[start][k] != -1 {
//			return memory[start][k]
//		}
//
//		result := math.MaxInt64
//
//		if k == 1 {
//			for i := start; i < n; i++ {
//				result = min(result, nums[i])
//			}
//		} else {
//			for i := start; i+2 < n; i++ {
//				minV := dfs(i+2, k-1)
//				result = min(result, max(minV, nums[i]))
//			}
//		}
//
//		memory[start][k] = result
//		return result
//	}
//
//	return dfs(0, k)
//}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
