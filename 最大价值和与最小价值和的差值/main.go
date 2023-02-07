package main

import "fmt"

func main() {
	fmt.Println(maxOutput(6, [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}, {3, 5}}, []int{9, 8, 7, 6, 10, 5}))
}

func maxOutput(n int, edges [][]int, price []int) int64 {
	neighbors := make([][]int, n)
	for _, edge := range edges {
		neighbors[edge[0]] = append(neighbors[edge[0]], edge[1])
		neighbors[edge[1]] = append(neighbors[edge[1]], edge[0])
	}

	result := int64(0)
	// 返回 带上端点的最大路径和, 不带端点的最大路径和
	var dfs func(root, father int) (int64, int64)
	dfs = func(root, father int) (int64, int64) {
		maxS1, maxS2 := int64(price[root]), int64(0)
		for _, neighbor := range neighbors[root] {
			if neighbor == father {
				continue
			}
			s1, s2 := dfs(neighbor, root)
			result = max(result, max(s1+maxS2, s2+maxS1))
			maxS1 = max(maxS1, s1+int64(price[root]))
			maxS2 = max(maxS2, s2+int64(price[root]))
		}

		return maxS1, maxS2
	}
	dfs(0, -1)

	return result
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
