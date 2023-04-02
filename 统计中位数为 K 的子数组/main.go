package main

import "fmt"

func main() {
	fmt.Println(countSubarrays([]int{3, 2, 1, 4, 5}, 4))
	fmt.Println(countSubarrays([]int{2, 3, 1}, 3))

}

func countSubarrays(nums []int, k int) int {

	n := len(nums)
	sums := make([]int, n+1)
	for i, num := range nums {
		sums[i+1] = sums[i]
		if num > k {
			sums[i+1]++
		}

		if num < k {
			sums[i+1]--
		}
	}
	//fmt.Println(sums)

	m := make(map[int]int)
	needCount := false
	ans := 0
	for i, sum := range sums {

		if i >= 1 && nums[i-1] == k {
			needCount = true
		}

		if needCount {
			//fmt.Println(i, sum, m[sum]+m[sum-1])
			ans += m[sum] + m[sum-1]
		} else {
			m[sum]++
		}
	}

	return ans
}
