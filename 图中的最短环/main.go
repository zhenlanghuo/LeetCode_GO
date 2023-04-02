package main

import "math"

func main() {

}

func findShortestCycle(n int, edges [][]int) int {
	g := make([][]int, n)
	for i := 0; i < len(edges); i++ {
		x, y := edges[i][0], edges[i][1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	dis := make([]int, n)
	ans := math.MaxInt64

	type pair struct {
		x  int
		fa int
	}
next:
	for start := 0; start < n; start++ {
		for i := 0; i < n; i++ {
			dis[i] = -1
		}
		queue := []*pair{{x: start, fa: -1}}
		dis[start] = 0
		for len(queue) != 0 {
			q := queue[0]
			queue = queue[1:]
			for _, y := range g[q.x] {
				if y == q.fa {
					continue
				}
				if dis[y] == -1 {
					dis[y] = dis[q.x] + 1
					queue = append(queue, &pair{x: y, fa: q.x})
				} else {
					ans = min(ans, dis[y]+dis[q.x]+1)
					continue next
				}
			}
		}
	}

	if ans == math.MaxInt64 {
		return -1
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
