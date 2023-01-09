package main

import "fmt"

func main() {
	fmt.Println(smallestValue(100))
}

func smallestValue(n int) int {
	primes := getPrime(n)
	fmt.Println(primes)
	originN := n
	for n > 0 {
		factors := primeFactor(n, primes)
		if len(factors) == 1 {
			return factors[0]
		}
		//fmt.Println(factors)
		temp := 0
		for _, factor := range factors {
			temp += factor
		}
		n = temp
		if n == originN {
			return n
		}
	}

	return 0
}

func primeFactor(n int, primes []int) []int {
	result := make([]int, 0)
	for n > 0 {
		//fmt.Println(n)
		for _, prime := range primes {
			if prime > n {
				break
			}
			if prime == n {
				result = append(result, n)
				return result
			}
			if n%prime == 0 {
				result = append(result, prime)
				n = n / prime
				break
			}
		}
	}
	return result
}

func getPrime(n int) []int {
	primes := make([]bool, n+1)
	for i := 0; i <= n; i++ {
		primes[i] = true
	}
	result := make([]int, 0)
	for i, prime := range primes {
		if i <= 1 {
			continue
		}
		if !prime {
			continue
		}
		if primes[i] {
			for j := i*i; j <= n; j += i {
				primes[j] = false
			}
			result = append(result, i)
		}

	}

	return result
}

//func getPrime(n int) []int {
//	primes := []int{}
//	isPrime := make([]bool, n+1)
//	for i := range isPrime {
//		isPrime[i] = true
//	}
//	for i := 2; i <= n; i++ {
//		if isPrime[i] {
//			primes = append(primes, i)
//		}
//		for _, p := range primes {
//			if i*p >= n {
//				break
//			}
//			isPrime[i*p] = false
//			if i%p == 0 {
//				break
//			}
//		}
//	}
//	return primes
//}
