package main

import "fmt"

func main() {
	fmt.Println(magicalString(20))
}

func magicalString(n int) int {
	if n <= 0 {
		return 0
	}

	if n <= 3 {
		return 1
	}

	fast := 2
	slow := 2
	array := make([]int, n+1)
	array[0] = 1
	array[1] = 2
	array[2] = 2
	for fast < n-1 {
		step := array[slow]
		next := 1
		if array[fast] == 1 {
			next = 2
		}
		for i := 0; i < step; i++ {
			array[fast+i+1] = next
		}
		fast += step
		slow++
	}

	//fmt.Println(array)

	count := 0
	for i := 0; i < n; i++ {
		if array[i] == 1 {
			count++
		}
	}

	return count
}
