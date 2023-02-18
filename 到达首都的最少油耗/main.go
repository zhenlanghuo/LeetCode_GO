package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumFuelCost([][]int{{3, 1}, {3, 2}, {1, 0}, {0, 4}, {0, 5}, {4, 6}}, 2))
	fmt.Println(minimumFuelCost([][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}}, 5))
}

type Node struct {
	p        int
	next     int
	distance int
	peoples  int
}

func minimumFuelCost(roads [][]int, seats int) int64 {
	n := len(roads) + 1
	g := make([][]int, n)
	for _, road := range roads {
		g[road[0]] = append(g[road[0]], road[1])
		g[road[1]] = append(g[road[1]], road[0])
	}

	var dfs func(x, fa int) int
	ans := int64(0)
	dfs = func(x, fa int) int {
		size := 1
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			temp := dfs(y, x)
			size += temp
			if temp%seats == 0 {
				ans += int64(temp / seats)
			} else {
				ans += int64(temp/seats) + 1
			}
		}

		return size
	}
	dfs(0, -1)

	return ans
}
