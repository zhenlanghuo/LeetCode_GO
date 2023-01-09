	package main

import "fmt"

func main() {
	fmt.Println(numIslands([][]byte{{1, 1, 1}, {0, 1, 0}, {1, 1, 1}}))
}

func numIslands(grid [][]byte) int {

	rowLen := len(grid)
	colLen := len(grid[0])

	byte0 := []byte("0")[0]
	byte1 := []byte("1")[0]

	byte0 = 0
	byte1 = 1

	result := 0

	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if grid[i][j] == byte0 {
				continue
			}
			result++
			setZero(grid, rowLen, colLen, i, j, byte0, byte1)
		}
	}

	return result
}

func setZero(grid [][]byte, rowLen, colLen, row, col int, byte0, byte1 byte) {
	grid[row][col] = byte0

	if row+1 < rowLen && grid[row+1][col] == byte1 {
		setZero(grid, rowLen, colLen, row+1, col, byte0, byte1)
	}

	if col+1 < colLen && grid[row][col+1] == byte1  {
		setZero(grid, rowLen, colLen, row, col+1, byte0, byte1)
	}

	if row-1 >=0 && grid[row-1][col] == byte1 {
		setZero(grid, rowLen, colLen, row-1, col, byte0, byte1)
	}

	if col-1 >= 0 && grid[row][col-1] == byte1  {
		setZero(grid, rowLen, colLen, row, col-1, byte0, byte1)
	}
}
