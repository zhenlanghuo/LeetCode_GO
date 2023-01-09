package main

import "fmt"

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
	fmt.Println(board)
}

func solveSudoku(board [][]byte) {
	rowUsed := make([][]bool, 9)
	colUsed := make([][]bool, 9)
	boxUsed := make([][]bool, 9)
	for i := 0; i < 9; i++ {
		rowUsed[i] = make([]bool, 9)
		colUsed[i] = make([]bool, 9)
		boxUsed[i] = make([]bool, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			b := board[i][j]
			if b != '.' {
				rowUsed[i][b-'0'-1] = true
				colUsed[b-'0'-1][j] = true
				boxUsed[(i/3)*3+j/3][b-'0'-1] = true
			}
		}
	}

	fmt.Println(rowUsed)
	fmt.Println(colUsed)
	fmt.Println(boxUsed)

	var backtrack func(i, j int) bool
	backtrack = func(i, j int) bool {
		if i >= 9 {
			return true
		}
		if board[i][j] != '.' {
			if j < 8 {
				return backtrack(i, j+1)
			} else {
				return backtrack(i+1, 0)
			}
		}

		for k := 0; k < 9; k++ {
			if rowUsed[i][k] {
				continue
			}
			if colUsed[k][j] {
				continue
			}
			if boxUsed[(i/3)*3+j/3][k] {
				continue
			}
			rowUsed[i][k] = true
			colUsed[k][j] = true
			boxUsed[(i/3)*3+j/3][k] = true
			board[i][j] = '0' + byte(k+1)

			//fmt.Println(i, j, k)

			if j < 8 {
				if backtrack(i, j+1) {
					return true
				}
			} else {
				if backtrack(i+1, 0) {
					return true
				}
			}

			rowUsed[i][k] = false
			colUsed[k][j] = false
			boxUsed[(i/3)*3+j/3][k] = false
			board[i][j] = '.'
		}

		return false
	}

	backtrack(0, 0)
}
