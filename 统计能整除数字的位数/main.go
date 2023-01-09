package main

import "fmt"

func main() {
	fmt.Println(countDigits(0))
}

func countDigits(num int) int {
	temp := num

	count := 0
	for temp > 0 {
		if num%(temp%10) == 0 {
			count++
		}
		temp = temp / 10
	}

	return count
}
