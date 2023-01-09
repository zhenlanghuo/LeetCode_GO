package main

import "fmt"

func main() {
	//fmt.Println(closestPrimes(10,19))
	//fmt.Println(closestPrimes(4,6))
	fmt.Println(closestPrimes(999998, 1000000))
}

func closestPrimes(left int, right int) []int {

	primes := calPrimes(right)

	fmt.Println(len(primes))

	for i, prime := range primes {
		if prime >= left {
			primes = primes[i:]
			break
		}

		if i == len(primes)-1 {
			return []int{-1, -1}
		}
	}

	fmt.Println(len(primes))

	//fmt.Println(primes)

	if len(primes) <= 1 {
		return []int{-1, -1}
	}

	result := []int{primes[0], primes[1]}
	for i := 1; i+1 < len(primes); i++ {
		if result[1]-result[0] > primes[i+1]-primes[i] {
			result[0], result[1] = primes[i], primes[i+1]
		}
	}

	return result
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
