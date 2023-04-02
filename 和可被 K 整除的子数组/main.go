package main

import "fmt"

func main() {

}

func subarraysDivByK(nums []int, k int) int {

	n := len(nums)
	sums := make([]int, n+1)
	for i, num := range nums {
		sums[i+1] = sums[i] + num
	}
	fmt.Println(sums)

	modMap := make(map[int]int)
	ans := 0
	for _, sum := range sums {
		modulus := (sum%k + k) % k
		ans += modMap[modulus]
		modMap[modulus]++
	}

	return ans
}
