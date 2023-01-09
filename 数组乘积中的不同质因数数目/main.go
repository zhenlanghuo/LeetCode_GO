package main

import "fmt"

func main() {
	fmt.Println(distinctPrimeFactors([]int{2,4,3,7,10,6}))
	fmt.Println(distinctPrimeFactors([]int{2,4,8,16}))
}

func distinctPrimeFactors(nums []int) int {

	primes := calPrimes(1000)
	fmt.Println(primes)

	resultMap := make(map[int]bool)

	for _, num := range nums {
		for _, prime := range primes {
			if num == 0 {
				break
			}

			if num%prime != 0 {
				continue
			}

			resultMap[prime] = true
			for num > 0 && num%prime == 0 {
				num = num / prime
			}
		}
	}

	return len(resultMap)
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
