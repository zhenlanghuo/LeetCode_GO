package main

import "fmt"

func main() {
	fmt.Println(minIncrements(7, []int{1, 5, 2, 2, 3, 3, 1}))
}

func minIncrements(n int, cost []int) int {
	ans := 0

	nums := make([]int, n)
	nums[0] = cost[0]
	for i := 1; i < n; i++ {
		nums[i] = nums[(i-1)/2] + cost[i]
	}
	fmt.Println(nums)

	nums = nums[n/2:]
	mx := 0
	for _, num := range nums {
		mx = max(mx, num)
	}
	fmt.Println(nums, mx)

	var dfs func(start, end int)
	dfs = func(start, end int) {
		//if start > end {
		//	return
		//}
		tempMx := 0
		for i := start; i <= end; i++ {
			//if nums[i] == mx {
			//	continue
			//}
			tempMx = max(tempMx, nums[i])
		}
		ans += mx - tempMx
		for i := start; i <= end; i++ {
			if nums[i] == mx {
				continue
			}
			nums[i] += mx - tempMx
		}
		//fmt.Println(start, end, tempMx, ans, nums)
		if start == end {
			return
		}
		dfs(start, (start+end)/2)
		dfs((start+end)/2+1, end)
	}
	dfs(0, len(nums)-1)

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
