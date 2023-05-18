package main

import "fmt"

func main() {
	//fmt.Println(distinctDifferenceArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(distinctDifferenceArray([]int{3, 2, 3, 4, 2}))
}

func distinctDifferenceArray(nums []int) []int {
	n := len(nums)
	m1 := make(map[int]int)
	m2 := make(map[int]int)

	ans := make([]int, n)

	diff1 := 0
	diff2 := 0
	for i := n - 1; i >= 0; i-- {
		m2[nums[i]]++
		if m2[nums[i]] == 1 {
			diff2++
		}

		//if m2[nums[i]] == 2 {
		//	diff2--
		//}
	}

	fmt.Println(m2, diff2)

	for i := 0; i < n; i++ {
		m1[nums[i]]++
		m2[nums[i]]--

		if m2[nums[i]] == 0 {
			diff2--
		}

		if m1[nums[i]] == 1 {
			diff1++
		}

		//fmt.Println(m1, m2, diff1, diff2)

		ans[i] = diff1 - diff2
	}

	return ans
}
