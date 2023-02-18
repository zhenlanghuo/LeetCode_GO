package main

import "fmt"

func main() {
	fmt.Println(unequalTriplets([]int{4, 4, 2, 4, 3}))
	fmt.Println(unequalTriplets([]int{1, 1, 1, 1, 1}))
}

func unequalTriplets(nums []int) int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}

	uqnums := make([]int, 0, len(numMap))
	for num, _ := range numMap {
		uqnums = append(uqnums, num)
	}

	ans := 0
	n := len(uqnums)
	for i, _ := range uqnums {
		for j := i + 1; j+1 < n; j++ {
			for k := j + 1; k < n; k++ {
				ans += numMap[uqnums[i]] * numMap[uqnums[j]] * numMap[uqnums[k]]
			}
		}
	}

	return ans
}
