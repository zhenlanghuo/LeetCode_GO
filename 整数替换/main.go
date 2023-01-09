package main

import "fmt"

func main() {
	fmt.Println(integerReplacement(100000000))
}

var memo map[int]int

func _integerReplacement(n int) (res int) {
	if n == 1 {
		return 0
	}
	if res, ok := memo[n]; ok {
		return res
	}
	defer func() { memo[n] = res }()
	if n%2 == 0 {
		return 1 + _integerReplacement(n/2)
	}
	return 2 + min(_integerReplacement(n/2), _integerReplacement(n/2+1))
}

func integerReplacement(n int) (res int) {
	memo = map[int]int{}
	return _integerReplacement(n)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}