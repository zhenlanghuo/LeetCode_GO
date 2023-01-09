package main

import "fmt"

type operateFunc func(grid [][]int, row, col int) (operateFunc, int, int)

func main() {
	grid := [][]int{{1,1,1,1,1,1},{-1,-1,-1,-1,-1,-1},{1,1,1,1,1,1},{-1,-1,-1,-1,-1,-1}}
	fmt.Printf("%v", findBall(grid))
}

func findBall(grid [][]int) []int {
	row := len(grid)
	col := len(grid[0])
	result := make([]int, col)
	for i:=0;i<col;i++ {
		of := down
		nextRow := 0
		nextCol := i
		for {
			fmt.Printf("i=%v, nextRow=%v, nextCol=%v\n", i, nextRow, nextCol)
			of, nextRow, nextCol = of(grid, nextRow, nextCol)
			if nextRow == -1 && nextCol == -1 {
				result[i] = -1
				break
			}
			if nextRow >= row {
				result[i] = nextCol
				break
			}
			if nextCol < 0 || nextCol >= col {
				result[i] = -1
				break
			}
		}
	}
	return result
}

func down(grid [][]int, row, col int) (operateFunc, int, int) {
	if grid[row][col] == 1 {
		return right, row, col + 1
	}

	if grid[row][col] == -1 {
		return left, row, col - 1
	}

	return down, -1, -1
}

func right(grid [][]int, row, col int) (operateFunc, int, int) {
	if grid[row][col] == -1 {
		return right, -1, -1
	}

	if grid[row][col] == 1 {
		return down, row + 1, col
	}

	return right, -1, -1
}

func left(grid [][]int, row, col int) (operateFunc, int, int) {
	if grid[row][col] == 1 {
		return left, -1, -1
	}

	if grid[row][col] == -1 {
		return down, row + 1, col
	}

	return left, -1, -1
}
