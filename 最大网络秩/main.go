package main

import (
	"fmt"
)

func main() {
	//fmt.Println(maximalNetworkRank(4, [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}}))
	//fmt.Println(maximalNetworkRank(5, [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}, {2, 3}, {2, 4}}))
	//fmt.Println(maximalNetworkRank(8, [][]int{{0, 1}, {1, 2}, {2, 3}, {2, 4}, {5, 6}, {5, 7}}))
	fmt.Println(maximalNetworkRank(2, [][]int{{1, 0}}))
}

func maximalNetworkRank(n int, roads [][]int) int {
	g := make([]map[int]struct{}, n)
	for i := 0; i < n; i++ {
		g[i] = make(map[int]struct{})
	}
	for _, road := range roads {
		x, y := road[0], road[1]
		g[x][y] = struct{}{}
		g[y][x] = struct{}{}
	}

	first, second := -1, -1
	for i := 0; i < n; i++ {
		if len(g[i]) > first {
			second = first
			first = len(g[i])
		} else if len(g[i]) != first && len(g[i]) > second {
			second = len(g[i])
		}
	}

	//fmt.Println(first, second)

	firstArr, secondArr := make([]int, 0), make([]int, 0)
	for i := 0; i < n; i++ {
		if first == len(g[i]) {
			firstArr = append(firstArr, i)
		}

		if second == len(g[i]) {
			secondArr = append(secondArr, i)
		}
	}

	//fmt.Println(firstArr, secondArr)

	ans := 0
	if len(firstArr) == 1 {
		for _, secondIndex := range secondArr {
			temp := first + second
			if _, ok := g[firstArr[0]][secondIndex]; ok {
				temp--
			}
			ans = max(ans, temp)
		}
	} else {
		x := len(firstArr)
		if x*(x-1)/2 > len(roads) {
			return 2 * first
		}
		for i := 0; i < x; i++ {
			for j := i + 1; j < x; j++ {
				temp := 2 * first
				if _, ok := g[firstArr[i]][firstArr[j]]; ok {
					temp--
				}
				ans = max(ans, temp)
			}
		}
	}

	return ans
}

//func maximalNetworkRank(n int, roads [][]int) int {
//	g := make([]map[int]struct{}, n)
//	for i := 0; i < n; i++ {
//		g[i] = make(map[int]struct{})
//	}
//	for _, road := range roads {
//		x, y := road[0], road[1]
//		g[x][y] = struct{}{}
//		g[y][x] = struct{}{}
//	}
//
//	ans := 0
//	for i := 0; i < n; i++ {
//		for j := i + 1; j < n; j++ {
//			temp := len(g[i]) + len(g[j])
//			if _, ok := g[i][j]; ok {
//				temp--
//			}
//			ans = max(ans, temp)
//		}
//	}
//
//	return ans
//}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
