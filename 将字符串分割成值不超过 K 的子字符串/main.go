package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(minimumPartition("165462", 60))
	//fmt.Println(minimumPartition("238182", 5))
	fmt.Println(minimumPartition("16839494817952275683974113656464314693", 169))
}

func minimumPartition(s string, k int) int {

	memory := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		memory[i] = -1
	}

	var dfs func(start int) int

	dfs = func(start int) int {
		if start >= len(s) {
			return 0
		}

		if memory[start] != -1 {
			return memory[start]
		}

		temp := 0
		result := math.MaxInt32
		for i := start; i < len(s); i++ {
			temp = temp*10 + int(s[i]-'0')
			if temp > k {
				break
			}
			result = min(result, dfs(i+1)+1)
		}
		memory[start] = result

		return result
	}

	result := dfs(0)
	if result == math.MaxInt32 {
		return -1
	}

	return result

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
