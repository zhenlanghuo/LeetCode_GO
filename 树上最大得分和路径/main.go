package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(mostProfitablePath([][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}, 3, []int{-2, 4, 2, -4, 6}))
	fmt.Println(mostProfitablePath([][]int{{0, 1}, {1, 2}, {2, 3}}, 3, []int{-5644, -6018, 1188, -8502}))
}

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(edges) + 1
	g := make([][]int, n)
	for _, edge := range edges {
		g[edge[0]] = append(g[edge[0]], edge[1])
		g[edge[1]] = append(g[edge[1]], edge[0])
	}
	node := make([]int, n)

	var findBob func(x, fa int) bool
	findBob = func(x, fa int) bool {
		if x == bob {
			node[x] = 1
			return true
		}

		for _, p := range g[x] {
			if p == fa {
				continue
			}
			if findBob(p, x) {
				node[x] = node[p] + 1
				return true
			}
		}

		return false
	}
	findBob(0, -1)
	fmt.Println(node)

	ans := math.MinInt64
	var dfs func(x, fa, time, point int)
	dfs = func(x, fa, time, point int) {
		temp := point
		if node[x] != 0 {
			if time == node[x]-1 {
				temp += amount[x] / 2
			} else if time < node[x]-1 {
				temp += amount[x]
			}
		} else {
			temp += amount[x]
		}

		//fmt.Println(x, fa, time, temp)
		if len(g[x]) == 1 && g[x][0] == fa {
			ans = max(ans, temp)
		}

		for _, p := range g[x] {
			if p == fa {
				continue
			}
			dfs(p, x, time+1, temp)
		}
	}
	dfs(0, -1, 0, 0)

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
