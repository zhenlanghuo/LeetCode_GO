package main

import "fmt"

func main() {
	//fmt.Println(minSubarray([]int{3, 1, 4, 2}, 6))
	//fmt.Println(minSubarray([]int{6, 3, 5, 2}, 9))
	//fmt.Println(minSubarray([]int{1, 2, 3}, 3))
	//fmt.Println(minSubarray([]int{1, 2, 3}, 7))
	//fmt.Println(minSubarray([]int{1000000000, 1000000000, 1000000000}, 3))
	//
	//fmt.Println(minSubarray([]int{4, 5, 8, 5, 4}, 7))
	fmt.Println(minSubarray([]int{8, 32, 31, 18, 34, 20, 21, 13, 1, 27, 23, 22, 11, 15, 30, 4, 2}, 148))
}

func minSubarray(nums []int, p int) int {

	n := len(nums)
	sums := make([]int, n+1)
	for i := 0; i < n; i++ {
		sums[i+1] = sums[i] + nums[i]
	}

	rem := sums[n] % p
	if rem == 0 {
		return 0
	}

	m := make(map[int]int)
	m[0] = -1
	ans := n
	for i := 0; i < n; i++ {
		temp := sums[i+1] % p
		if v, ok := m[(sums[i+1]-rem)%p]; ok {
			ans = min(ans, i-v)
		}
		m[temp] = i
	}

	if ans == n {
		ans = -1
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
