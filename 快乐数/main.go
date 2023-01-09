package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("%v", isHappy(2))
}

func isHappy(n int) bool {
	array := make([]int, 1000)
	next := n
	for {
		next = caculate(next)
		fmt.Printf("next=%v\n", next)
		if next == 1 {
			return true
		}
		if array[next] == 1 {
			return false
		}
		array[next] = 1
	}
}

func caculate(n int) int {
	result := 0
	nStr := strconv.FormatInt(int64(n), 10)
	for i := 0; i < len(nStr); i++ {
		item, _ := strconv.ParseUint(string(nStr[i]), 10, 64)
		result += int(item * item)
	}

	return result
}
