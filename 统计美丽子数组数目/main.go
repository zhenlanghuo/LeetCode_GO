package main

import "fmt"

func main() {
	fmt.Println(beautifulSubarrays([]int{0, 0}))
}

func beautifulSubarrays(nums []int) int64 {
	n := len(nums)
	xorSums := make([]int, n+1)
	for i, num := range nums {
		xorSums[i+1] = xorSums[i] ^ num
	}

	xorSumMap := make(map[int]int)
	ans := int64(0)
	for _, xorSum := range xorSums {
		ans += int64(xorSumMap[xorSum])
		xorSumMap[xorSum]++
	}

	return ans
}
