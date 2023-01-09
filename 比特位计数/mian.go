package main

import (
	"fmt"
)

func main() {
	fmt.Println(countBits(15))
}

func countBits(n int) []int {

	if n <= 1 {
		switch n {
		case 0:
			return []int{0}
		case 1:
			return []int{0, 1}
		}
	}

	result := make([]int, n+1)
	result[0] = 0
	result[1] = 1

	temp := 2
	i := 0
	for i+temp <= n {
		result[i+temp] = result[i] + 1
		if i == temp-1 {
			i = 0
			temp = temp * 2
		} else {
			i++
		}

	}

	return result
}
