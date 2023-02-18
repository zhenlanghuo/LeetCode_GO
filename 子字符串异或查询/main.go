package main

import (
	"fmt"
)

func main() {
	fmt.Println(substringXorQueries("101101", [][]int{{0, 5}, {1, 2}}))
}

func substringXorQueries(s string, queries [][]int) [][]int {
	result := make([][]int, 0, len(queries))
	memory := make(map[int][]int, 0)
	for i := 0; i < len(s); i++ {
		x := 0
		for j := i; j-i+1 <= 30 && j < len(s); j++ {
			x = x*2 + int(s[j]-'0')
			v, ok := memory[x]
			if !ok {
				memory[x] = []int{i, j}
			} else if v[1]-v[0] > j-i {
				memory[x][0], memory[x][1] = i, j
			}
		}
	}

	notFound := []int{-1, -1}
	for _, query := range queries {
		v, ok := memory[query[0]^query[1]]
		if ok {
			result = append(result, v)
		} else {
			result = append(result, notFound)
		}
	}

	return result
}
