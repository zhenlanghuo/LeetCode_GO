package main

import (
	"fmt"
	"math"
)

func main() {
	solve(6, [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 4}, {2, 5}})
}

func solve(n int, edges [][]int) []int {
	dis := make([][]int, n)
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		dis[i] = make([]int, n)
	}

	for _, edge := range edges {
		dis[edge[0]][edge[1]] = 1
		dis[edge[1]][edge[0]] = 1
	}

	var find func(i, j int)
	find = func(i, j int) {
		if dis[i][j] != 0 {
			return
		}

		result := math.MaxInt32
		index := -1
		for k := 0; k < n; k++ {
			if k == i || k == j {
				continue
			}
			if used[k] {
				continue
			}
			if dis[i][k] == 0 {
				continue
			}
			used[k] = true
			find(j, k)
			if dis[j][k] == 0 {
				continue
			}
			if dis[i][k]+dis[j][k] < result {
				//fmt.Println(i, j, k)
				result = dis[i][k] + dis[j][k]
				index = k
			}
			used[k] = false
		}
		if index == -1 {
			return
		}
		//fmt.Println(i, j, index, dis[i][index], dis[j][index], dis)
		dis[i][j] = dis[i][index] + dis[j][index]
		dis[j][i] = dis[i][index] + dis[j][index]
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			find(i, j)
		}
	}

	fmt.Println(dis)

	return nil
}
