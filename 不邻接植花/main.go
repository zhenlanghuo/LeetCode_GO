package main

import "fmt"

func main() {
	fmt.Println(gardenNoAdj(3, [][]int{{1, 2}, {2, 3}, {3, 1}}))
}

func gardenNoAdj(n int, paths [][]int) []int {
	ans := make([]int, n+1)
	dir := make([][]int, n+1)
	for _, path := range paths {
		x, y := path[0], path[1]
		dir[x] = append(dir[x], y)
		dir[y] = append(dir[y], x)
	}

	neighborFlower := make([]bool, 5)

	var dfs func(x int)
	dfs = func(x int) {
		for i := 0; i < len(neighborFlower); i++ {
			neighborFlower[i] = false
		}
		for _, y := range dir[x] {
			neighborFlower[ans[y]] = true
		}
		//fmt.Println(x, neighborFlower)
		for i := 1; i < len(neighborFlower); i++ {
			if !neighborFlower[i] {
				ans[x] = i
				break
			}
		}
		for _, y := range dir[x] {
			if ans[y] != 0 {
				continue
			}
			dfs(y)
		}
	}
	for i := 1; i < len(ans); i++ {
		if ans[i] == 0 {
			dfs(i)
		}
	}
	return ans[1:]
}
