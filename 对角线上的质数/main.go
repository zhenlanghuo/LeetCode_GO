package main

import "fmt"

func main() {
	fmt.Println(isPrime(1))
	fmt.Println(diagonalPrime([][]int{}))
}

func diagonalPrime(nums [][]int) int {
	n := len(nums)
	mx := 0
	for i := 0; i < n; i++ {
		if isPrime(nums[i][i]) && nums[i][i] > mx {
			mx = nums[i][i]
		}
		if isPrime(nums[i][n-i-1]) && nums[i][n-i-1] > mx {
			mx = nums[i][n-i-1]
		}
	}

	return mx
}

func isPrime(num int) bool {
	if num == 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
