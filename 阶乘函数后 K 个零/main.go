package main

import "fmt"

func main() {
	k := 5
	fmt.Println(preimageSizeFZF(k))
	fmt.Println(findMax(k))
	fmt.Println(findMin(k))
}

func preimageSizeFZF(k int) int {
	return findMax(k) - findMin(k) + 1
}

func trailingZeroes(n int) int {
	result := 0
	for n/5 != 0 {
		result += n / 5
		n = n / 5
	}

	return result
}

func findMin(k int) int {
	left, right := 0, 5*k+1
	for left < right {
		mid := (left + right) / 2
		if trailingZeroes(mid) >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func findMax(k int) int {
	left, right := 0, 5*(k+1)+1
	for left < right {
		mid := (left + right) / 2
		if trailingZeroes(mid) > k {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left - 1
}
