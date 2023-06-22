package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxIncreasingCells([][]int{{3, 1}, {3, 4}}))
	//fmt.Println(maxIncreasingCells([][]int{{1, 1}, {1, 1}}))
	fmt.Println(maxIncreasingCells([][]int{{1, 1}, {1, 1}}))
	fmt.Println(maxIncreasingCells([][]int{{3, 1, 6}, {-9, 5, 7}}))
}

func maxIncreasingCells(mat [][]int) int {

	type Node struct {
		val  int
		i, j int
	}

	n, m := len(mat), len(mat[0])
	RowNodes := make([][]Node, n)
	for i := 0; i < n; i++ {
		RowNodes[i] = make([]Node, m)
		for j := 0; j < m; j++ {
			RowNodes[i][j] = Node{
				val: mat[i][j],
				i:   i,
				j:   j,
			}
		}
		sort.Slice(RowNodes[i], func(a, b int) bool {
			if RowNodes[i][a].val == RowNodes[i][b].val {
				return RowNodes[i][a].j < RowNodes[i][b].j
			}
			return RowNodes[i][a].val < RowNodes[i][b].val
		})
	}
	ColNodes := make([][]Node, m)
	for j := 0; j < m; j++ {
		ColNodes[j] = make([]Node, n)
		for i := 0; i < n; i++ {
			ColNodes[j][i] = Node{
				val: mat[i][j],
				i:   i,
				j:   j,
			}
		}
		sort.Slice(ColNodes[j], func(a, b int) bool {
			if ColNodes[j][a].val == ColNodes[j][b].val {
				return ColNodes[j][a].i < ColNodes[j][b].i
			}
			return ColNodes[j][a].val < ColNodes[j][b].val
		})
	}

	//fmt.Println(RowNodes)
	//fmt.Println(ColNodes)

	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, m)
		for j := 0; j < m; j++ {
			result[i][j] = -1
		}
	}
	ans := 0

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if result[i][j] != -1 {
			return result[i][j]
		}

		mx := 0
		colIndex := sort.Search(len(RowNodes[i]), func(a int) bool {
			return RowNodes[i][a].val > mat[i][j]
		})
		temp := 0
		if colIndex >= 0 && colIndex < len(RowNodes[i]) {
			temp = RowNodes[i][colIndex].val
		}
		for colIndex >= 0 && colIndex < len(RowNodes[i]) && RowNodes[i][colIndex].val == temp {
			mx = max(mx, dfs(RowNodes[i][colIndex].i, RowNodes[i][colIndex].j))
			colIndex++
		}
		//for k := len(RowNodes[i]) - 1; k >= colIndex; k-- {
		//	mx = max(mx, dfs(RowNodes[i][k].i, RowNodes[i][k].j))
		//}

		rowIndex := sort.Search(len(ColNodes[j]), func(a int) bool {
			return ColNodes[j][a].val > mat[i][j]
		})
		temp = 0
		if rowIndex >= 0 && rowIndex < len(ColNodes[j]) {
			temp = ColNodes[j][rowIndex].val
		}
		for rowIndex >= 0 && rowIndex < len(ColNodes[j]) && ColNodes[j][rowIndex].val == temp {
			mx = max(mx, dfs(ColNodes[j][rowIndex].i, ColNodes[j][rowIndex].j))
			rowIndex++
		}

		//for rowIndex < len(ColNodes[j]) {
		//	mx = max(mx, dfs(ColNodes[j][rowIndex].i, ColNodes[j][rowIndex].j))
		//	rowIndex++
		//}
		//for k := len(ColNodes[j]) - 1; k >= rowIndex; k-- {
		//	mx = max(mx, dfs(ColNodes[j][k].i, ColNodes[j][k].j))
		//}
		result[i][j] = mx + 1
		ans = max(ans, mx+1)
		return mx + 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dfs(i, j)
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
