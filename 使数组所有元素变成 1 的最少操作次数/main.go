package main

import "math"

func main() {

}

func minOperations(nums []int) int {

	n := len(nums)
	count1 := 0
	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			count1++
		}
	}

	if count1 > 0 {
		return n - count1
	}

	maxGcd := nums[0]
	for i := 1; i < n; i++ {
		maxGcd = gcd(maxGcd, nums[i])
	}

	if maxGcd != 1 {
		return -1
	}

	mn := math.MaxInt64
	for i := 0; i < n-1; i++ {
		maxGcd := nums[i]
		for j := i + 1; j < n; j++ {
			maxGcd = gcd(maxGcd, nums[j])
			if maxGcd == 1 {
				if mn > j-i-1 {
					mn = j - i - 1
				}
			}
		}
	}

	return n + mn
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
