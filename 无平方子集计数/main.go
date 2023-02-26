package main

import "fmt"

func main() {
	//fmt.Println(squareFreeSubsets([]int{3, 4, 4, 5}))
	fmt.Println(squareFreeSubsets([]int{1, 2, 6, 15, 7, 19, 6, 29, 28, 24, 21, 25, 25, 18, 9, 6, 20, 21, 8, 24, 14, 19, 24, 28, 30, 27, 13, 21, 1, 23, 13, 29, 24, 29, 18, 7}))
}

func squareFreeSubsets(nums []int) int {

	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	sum := 1
	for _, prime := range primes {
		sum *= prime
	}
	//fmt.Println(sum)

	m := make(map[int]int)
	ans := 0
	for _, num := range nums {
		factors := calFactor(num, primes)
		//fmt.Println(factors)
		flag := true
		for _, v := range factors {
			if v > 1 {
				flag = false
				break
			}
		}
		if !flag {
			continue
		}

		if num == 1 {
			for k := range m {
				m[k] = (m[k] + m[k]) % (1e9 + 7)
				ans = (ans + ans) % (1e9 + 7)
			}

			if m[num] == 0 {
				m[num] = 2
				ans = (ans + 1) % (1e9 + 7)
			}
		} else {
			m[num] = (m[num] + 1 + m[1] - 1) % (1e9 + 7)
			ans = (ans + 1 + m[1] - 1) % (1e9 + 7)
			for num1, _ := range m {
				factors1 := calFactor(num1, primes)
				flag1 := true
				for k := range factors1 {
					if _, ok := factors[k]; ok {
						flag1 = false
						break
					}
				}
				if flag1 {
					ans = (ans + m[num1] + m[1] - 1) % (1e9 + 7)
					m[num1*num] = (m[num1*num] + m[num1] + m[1] - 1) % (1e9 + 7)
				}
			}
		}

		//for _, count := range m {
		//	ans = (ans + count) % (1e9 + 7)
		//}
		fmt.Println(m, num, ans)
	}

	return ans
}

func calFactor(num int, primes []int) map[int]int {
	m := make(map[int]int)
	for _, prime := range primes {
		if num < prime {
			break
		}
		for num > 1 && num%prime == 0 {
			m[prime]++
			num = num / prime
		}
	}

	return m
}
