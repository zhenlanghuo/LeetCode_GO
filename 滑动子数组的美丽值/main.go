package main

import "fmt"

func main() {
	fmt.Println(getSubarrayBeauty([]int{1, -1, -3, -2, 3}, 3, 2))
}

func getSubarrayBeauty(nums []int, k int, x int) []int {

	m := make(map[int]int)
	n := len(nums)
	ans := make([]int, n-k+1)
	for i := 0; i < k; i++ {
		m[nums[i]]++
	}
	ans[0] = cal(m, x)
	//fmt.Println(m, ans)

	for i := k; i < n; i++ {
		m[nums[i-k]]--
		m[nums[i]]++
		ans[i-k+1] = cal(m, x)
		//fmt.Println(m, ans)
	}

	return ans
}

func cal(m map[int]int, x int) int {
	count := 0
	for i := -50; i <= -1; i++ {
		count += m[i]
		if count >= x {
			return i
		}
	}
	return 0
}
