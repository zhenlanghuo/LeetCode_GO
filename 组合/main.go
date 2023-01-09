package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
}

var (
	result [][]int
)

func combine(n int, k int) [][]int {
	result = make([][]int, 0)
	solve(1, k, n, make([]int, 0, 1<<uint64(n)))
	return result
}

func solve(num, k, n int, temp []int) {
	if len(temp) + n-num+1 < k {
		return
	}

	if num == n+1 || len(temp) == k {
		if len(temp) != k {
			return
		}
		//fmt.Println(temp)
		clone := make([]int, len(temp))
		copy(clone, temp)
		result = append(result, clone)
		return
	}

	solve(num+1, k, n, append(temp, num))
	solve(num+1, k, n, temp)
}
