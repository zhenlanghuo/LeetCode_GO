package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(primeSubOperation([]int{4, 9, 6, 10}))
	fmt.Println(primeSubOperation([]int{6, 8, 11, 12}))
	fmt.Println(primeSubOperation([]int{5, 8, 3}))
	fmt.Println(primeSubOperation([]int{2, 2}))
}

func primeSubOperation(nums []int) bool {

	primes := calPrimes(1000)
	//fmt.Println(primes)

	minNum := 0
	for _, num := range nums {
		if num <= minNum {
			return false
		}
		i := sort.SearchInts(primes, num-minNum) - 1
		//fmt.Println(num, i, minNum)

		temp := num
		if i >= 0 {
			temp = num - primes[i]
		}

		if temp <= minNum {
			return false
		}
		minNum = temp
		//fmt.Println(num, i, minNum)
	}

	return true
}

func calPrimes(n int) []int {
	isPrime := make([]bool, n+1)
	for i := 0; i <= n; i++ {
		isPrime[i] = true
	}

	result := make([]int, 0)
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			result = append(result, i)
		}
		for j := i * i; j <= n; j += i {
			isPrime[j] = false
		}
	}
	return result
}
