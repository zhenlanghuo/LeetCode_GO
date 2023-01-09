package main

import "fmt"

func main() {
	fmt.Println(edgeScore([]int{3, 3, 3, 0}))
}

func edgeScore(edges []int) int {
	n := len(edges)
	count := make([]int, n)
	max := 0
	result := 0
	for i, v := range edges {
		count[v] += i
		if count[v] > max {
			max = count[v]
			result = v
		}
		if count[v] == max && result > v {
			result = v
		}
	}
	return result
}
