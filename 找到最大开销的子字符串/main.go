package main

import "fmt"

func main() {
	fmt.Println(maximumCostSubstring("adaa", "d", []int{-1000}))
	fmt.Println(maximumCostSubstring("abc", "abc", []int{-1, -1, -1}))
}

func maximumCostSubstring(s string, chars string, vals []int) int {

	mChar := make(map[byte]int)

	for i, b := range chars {
		mChar[byte(b)] = vals[i]
	}

	n := len(s)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = int(s[i]-'a') + 1
		if v, ok := mChar[s[i]]; ok {
			nums[i] = v
		}
	}

	fmt.Println(nums)

	ans := 0
	cur := ans
	for i := 0; i < n; i++ {
		if cur+nums[i] < 0 {
			cur = 0
		} else {
			cur += nums[i]
			ans = max(ans, cur)
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
