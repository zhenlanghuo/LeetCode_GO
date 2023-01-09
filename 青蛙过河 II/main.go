package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxJump([]int{0, 4, 5, 9}))
}

func maxJump(stones []int) int {
	n := len(stones)

	result := math.MaxInt32

	var dfs func(start int, temp []int)
	dfs = func(start int, temp []int) {
		if start == n-1 {
			curMax := 0
			cur := 0
			tempMap := make(map[int]bool)
			for _, stoneIndex := range temp {
				tempMap[stoneIndex] = true
				curMax = max(abs(stones[cur]-stones[stoneIndex]), curMax)
				if curMax > result {
					return
				}
				cur = stoneIndex
			}
			curMax = max(abs(stones[cur]-stones[n-1]), curMax)
			if curMax > result {
				return
			}

			other := make([]int, 0, n-2-len(temp))
			for i := 1; i <= n-2; i++ {
				if _, ok := tempMap[i]; !ok {
					other = append(other, i)
				}
			}

			cur = 0
			for _, stoneIndex := range other {
				tempMap[stoneIndex] = true
				curMax = max(abs(stones[cur]-stones[stoneIndex]), curMax)
				if curMax > result {
					return
				}
				cur = stoneIndex
			}
			curMax = max(abs(stones[cur]-stones[n-1]), curMax)
			if curMax > result {
				return
			}

			result = min(curMax, result)
			return
		}
		dfs(start+1, append(temp, start))
		dfs(start+1, temp)
	}

	dfs(1, nil)

	return result

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
