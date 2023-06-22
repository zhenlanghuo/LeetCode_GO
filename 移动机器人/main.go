package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sumDistance([]int{-2, 0, 2}, "RLL", 3))
	fmt.Println(sumDistance([]int{1, 0}, "RL", 2))
}

func sumDistance(nums []int, s string, d int) int {

	n := len(nums)
	for i := 0; i < n; i++ {
		switch s[i] {
		case 'L':
			nums[i] -= d
		case 'R':
			nums[i] += d
		}
	}

	sort.Ints(nums)

	ans := 0
	mod := int(1e9 + 7)
	for i := 1; i < n; i++ {
		deta := nums[i] - nums[i-1]
		//ans = (ans%mod + (i*(n-i)*(deta%mod))%mod) % mod
		ans = (ans + i*(n-i)*deta) % mod
	}

	return ans
}
