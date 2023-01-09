package main

import "fmt"

func main() {
	//matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	//matrix := [][]int{{1, 2}, {3, 4}}
	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < len(matrix)/2; i++ {
		rotate_(matrix, i, n-1-i)
	}
}


func rotate_(matrix [][]int, start, end int) {
	n := end - start
	for i := 0; i < n; i++ {
		indexList := [][]int{{0 + start, i + start}, {i + start, n + start}, {n + start, n + start - i}, {n + start - i, 0 + start}}
		swap(matrix, indexList)
	}
}

func swap(matrix [][]int, indexList [][]int) {
	temp := matrix[indexList[0][0]][indexList[0][1]]
	matrix[indexList[0][0]][indexList[0][1]] = matrix[indexList[3][0]][indexList[3][1]]
	matrix[indexList[3][0]][indexList[3][1]] = matrix[indexList[2][0]][indexList[2][1]]
	matrix[indexList[2][0]][indexList[2][1]] = matrix[indexList[1][0]][indexList[1][1]]
	matrix[indexList[1][0]][indexList[1][1]] = temp
}
