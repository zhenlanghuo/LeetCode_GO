package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
}

func subarraySum(nums []int, k int) int {
	n := len(nums)
	sums := make([]int, n+1)
	for i, num := range nums {
		sums[i+1] = sums[i] + num
	}
	fmt.Println(sums)

	ans := 0
	sumMap := make(map[int]int)
	for _, sum := range sums {
		ans += sumMap[sum-k]
		sumMap[sum]++
	}

	return ans
}
