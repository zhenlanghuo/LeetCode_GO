package main

import "fmt"

func main() {
	fmt.Println(findSmallestInteger([]int{1, -10, 7, 13, 6, 8}, 5))
	fmt.Println(findSmallestInteger([]int{1, -10, 7, 13, 6, 8}, 7))
	fmt.Println(findSmallestInteger([]int{3, 0, 3, 2, 4, 2, 1, 1, 0, 4}, 5))
}

func findSmallestInteger(nums []int, value int) int {
	n := len(nums)
	mods := make([]bool, n+1)
	m := make(map[int]int)
	for _, num := range nums {
		mod := ((num % value) + value) % value
		if mod+value*m[mod] < n+1 {
			mods[mod+value*m[mod]] = true
		}
		m[mod]++
	}

	fmt.Println(mods)

	for i, v := range mods {
		if !v {
			return i
		}
	}

	return 0
}
