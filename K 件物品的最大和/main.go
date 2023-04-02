package main

import "fmt"

func main() {
	fmt.Println(kItemsWithMaximumSum(3, 2, 0, 2))
	fmt.Println(kItemsWithMaximumSum(3, 0, 2, 4))
}

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	ans := 0

	for k > 0 {
		if numOnes > 0 {
			temp := min(numOnes, k)
			ans += 1 * temp
			k -= temp
			numOnes -= temp
		}

		if numZeros > 0 {
			temp := min(numZeros, k)
			ans += 0 * temp
			k -= temp
			numZeros -= temp
		}

		if numNegOnes > 0 {
			temp := min(numNegOnes, k)
			ans -= 1 * temp
			k -= temp
			numNegOnes -= temp
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
