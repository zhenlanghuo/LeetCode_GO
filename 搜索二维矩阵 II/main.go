package main

import "fmt"

func main() {
	fmt.Println(search([]int{1, 2, 4}, 3))
}

//func searchMatrix(matrix [][]int, target int) bool {
//	row, col := len(matrix), len(matrix[0])
//
//	x, y := 0, col-1
//	for x < row && y >= 0 {
//		if matrix[x][y] == target {
//			return true
//		}
//
//		if matrix[x][y] < target {
//			x++
//		} else {
//			y--
//		}
//	}
//
//	return false
//}

func searchMatrix(matrix [][]int, target int) bool {

	for _, array := range matrix {
		if search(array, target) {
			return true
		}
	}

	return false
}

func search(array []int, target int) bool {
	if len(array) == 0 {
		return false
	}

	mid := len(array) / 2
	if array[mid] == target {
		return true
	}

	return search(array[:mid], target) || search(array[mid+1:], target)
}
