package main

import "fmt"

func main() {
	fmt.Println(circularGameLosers(5, 2))
}

func circularGameLosers(n int, k int) []int {
	count := make([]int, n)
	ans := make([]int, 0, n)

	i := 0
	temp := k
	for {
		//fmt.Println(i)
		count[i]++
		if count[i] == 2 {
			break
		}
		i = (i + temp) % n
		temp += k
	}

	for i, c := range count {
		if c == 0 {
			ans = append(ans, i+1)
		}
	}

	return ans
}
