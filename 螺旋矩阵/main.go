package main

import "fmt"

func main() {

	matrix := [][]int{{23,18,20,26,25},{24,22,3,4,4},{15,22,2,24,29},{18,15,23,28,28}}
	//matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Printf("%v", spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) []int {

	n := len(matrix)
	if n == 0 {
		return []int{}
	}
	m := len(matrix[0])

	result := make([]int, 0, len(matrix)*len(matrix[0]))
	resultMap := make(map[int]bool)

	operateFuncs := []func(n, m, row, col int) (int, int){right, down, left, up}

	i := 0
	nextRow := 0
	nextCol := 0
	for {
		result = append(result, matrix[nextRow][nextCol])
		resultMap[nextRow*m+nextCol] = true
		if len(result) == n*m {
			break
		}
		nextRow_ := -1
		nextCol_ := -1
		for {
			nextRow_, nextCol_ = operateFuncs[i](n, m, nextRow, nextCol)
			if (nextRow_ == -1 && nextCol_ == -1) || resultMap[nextRow_*m+nextCol_] == true {
				i = (i + 1) % len(operateFuncs)
				continue
			}
			break
		}
		nextRow = nextRow_
		nextCol = nextCol_
		fmt.Printf("%v, %v, %v \n", nextRow, nextCol, i)
	}

	return result
}

func right(n, m, row, col int) (int, int) {
	if row < n && col+1 < m {
		return row, col + 1
	}

	return -1, -1
}

func down(n, m, row, col int) (int, int) {
	if row+1 < n && col < m {
		return row + 1, col
	}

	return -1, -1
}

func left(n, m, row, col int) (int, int) {
	if row < n && col-1 >= 0 {
		return row, col - 1
	}

	return -1, -1
}

func up(n, m, row, col int) (int, int) {
	if row-1 >= 0 && col < m {
		return row - 1, col
	}

	return -1, -1
}
