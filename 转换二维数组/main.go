package main

import "fmt"

func main() {
	fmt.Println(findMatrix([]int{1, 3, 4, 1, 2, 3, 1}))
	fmt.Println(findMatrix([]int{1, 3, 4, 1, 2, 3, 1}))
}

func findMatrix(nums []int) [][]int {
	m := make(map[int]int)

	for _, num := range nums {
		m[num]++
	}
	fmt.Println(m)

	ans := make([][]int, 0)
	for len(m) != 0 {
		temp := make([]int, 0)
		for k, v := range m {
			temp = append(temp, k)
			if v == 1 {
				delete(m, k)
			} else {
				m[k]--
			}
		}
		//fmt.Println(m)
		ans = append(ans, temp)
	}
	return ans
}
