package main

import "fmt"

func main() {

}

func checkSubarraySum(nums []int, k int) bool {

	n := len(nums)
	sums := make([]int, n+1)
	for i, num := range nums {
		sums[i+1] = sums[i] + num
	}
	fmt.Println(sums)

	modMap := make(map[int]int)
	for i, sum := range sums {
		if v, ok := modMap[sum%k]; ok {
			if i-v >= 2 {
				return true
			}
		} else {
			modMap[sum%k] = i
		}
	}

	return false
}
