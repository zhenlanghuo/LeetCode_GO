package main

import "fmt"

func main() {
	fmt.Println(countSubgraphsForEachDiameter(4, [][]int{{1, 2}, {2, 3}, {2, 4}}))
}

func countSubgraphsForEachDiameter(n int, edges [][]int) []int {

	// 建树
	g := make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0]-1, edge[1]-1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	inSet, visited := [15]bool{}, [15]bool{}
	diameter := 0
	// 求直径
	var dfs func(x int) int
	dfs = func(x int) int {
		visited[x] = true
		maxLen := 0
		for _, child := range g[x] {
			if !visited[child] && inSet[child] {
				ml := dfs(child) + 1
				diameter = max(diameter, ml+maxLen)
				maxLen = max(maxLen, ml)
			}
		}
		return maxLen
	}

	// 求子集
	ans := make([]int, n-1)
	var f = func(start int) {}
	f = func(start int) {
		if start == n {
			for x, in := range inSet {
				if in {
					visited = [15]bool{}
					diameter = 0
					dfs(x)
					break
				}
			}
			if diameter > 0 && inSet == visited {
				ans[diameter-1]++
			}
			return
		}

		// 不选
		f(start + 1)
		// 选
		inSet[start] = true
		f(start + 1)
		inSet[start] = false
	}
	f(0)

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
