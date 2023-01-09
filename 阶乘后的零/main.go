package main

import "fmt"

func main() {
	fmt.Println(trailingZeroes(25))
}

func trailingZeroes(n int) int {
	result := 0
	for n/5 != 0 {
		result += n / 5
		n = n / 5
	}

	return result
}
