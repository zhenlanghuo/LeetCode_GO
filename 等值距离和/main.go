package main

import "fmt"

func main() {
	fmt.Println(distance([]int{1, 3, 1, 1, 2}))
}

func distance(nums []int) []int64 {

	n := len(nums)
	m := make(map[int][]int)
	for i, num := range nums {
		m[num] = append(m[num], i)
	}
	//fmt.Println(m)

	ans := make([]int64, n)
	for _, arr := range m {
		//fmt.Println(arr)
		for i := 0; i < len(arr); i++ {
			ans[arr[0]] += int64(abs(arr[i] - arr[0]))
		}
		//fmt.Println(ans[arr[0]])

		for i := 1; i < len(arr); i++ {
			ans[arr[i]] = ans[arr[i-1]] + int64(((i-1)-(len(arr)-1-i))*(arr[i]-arr[i-1]))
		}
	}

	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
