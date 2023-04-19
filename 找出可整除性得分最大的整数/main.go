package main

import "sort"

func main() {

}

func maxDivScore(nums []int, divisors []int) int {
	sort.Ints(divisors)
	ans := divisors[0]
	mx := 0

	for _, divisor := range divisors {
		sum := 0
		for _, num := range nums {
			if num%divisor == 0 {
				sum++
			}
		}
		if sum > mx {
			mx = sum
			ans = divisor
		}
	}
	return ans
}
