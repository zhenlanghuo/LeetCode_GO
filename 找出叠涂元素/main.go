package main

func main() {

}

func firstCompleteIndex(arr []int, mat [][]int) int {
	ans := int(1e5)

	arrMap := make(map[int]int)
	matMap := make(map[int][]int)
	n, m := len(mat), len(mat[0])

	for i, num := range arr {
		arrMap[num] = i
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			matMap[mat[i][j]] = []int{i, j}
		}
	}

	rowCount := make([]int, n)
	colCount := make([]int, m)

	for i, num := range arr {
		row, col := matMap[num][0], matMap[num][1]
		rowCount[row]++
		colCount[col]++
		if rowCount[row] == m || colCount[col] == n {
			return i
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
