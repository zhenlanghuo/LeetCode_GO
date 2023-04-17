package main

import (
	"fmt"
)

func main() {

}

func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	dir := make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		dir[x] = append(dir[x], y)
		dir[y] = append(dir[y], x)
	}

	var find func(x, fa, end int, result []int) ([]int, bool)
	find = func(x, fa, end int, result []int) ([]int, bool) {
		if x == end {
			result = append(result, x)
			temp := make([]int, len(result))
			copy(temp, result)
			return temp, true
		}

		for _, y := range dir[x] {
			if y == fa {
				continue
			}
			temp, ok := find(y, x, end, append(result, x))
			if ok {
				return temp, ok
			}
		}

		return nil, false
	}

	passMap := make(map[int]int)
	for _, trip := range trips {
		start, end := trip[0], trip[1]
		roads, _ := find(start, -1, end, nil)
		fmt.Println(start, end, roads)
		for _, road := range roads {
			passMap[road]++
		}
	}

	var dfs func(x, fa int) (int, int)
	dfs = func(x, fa int) (int, int) {
		halfResult := 0
		noHalfResult := 0

		// 不减半
		for _, y := range dir[x] {
			if y == fa {
				continue
			}
			r1, r2 := dfs(y, x)
			halfResult += r2
			noHalfResult += min(r1, r2)
		}
		halfResult += (price[x] / 2) * passMap[x]
		noHalfResult += price[x] * passMap[x]

		return halfResult, noHalfResult
	}

	r1, r2 := dfs(0, -1)

	return min(r1, r2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
