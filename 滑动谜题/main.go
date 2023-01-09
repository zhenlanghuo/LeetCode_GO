package main

import "fmt"

func main() {
	fmt.Println(slidingPuzzle([][]int{{1, 2, 3}, {4, 0, 5}}))
	fmt.Println(boardToInt([][]int{{1, 2, 3}, {4, 5, 0}}))
}

func slidingPuzzle(board [][]int) int {
	if boardToInt(board) == 123450 {
		return 0
	}

	n, m := len(board), len(board[0])
	visited := make(map[int]int)
	queue := make([][][]int, 0)
	queue = append(queue, board)
	visited[boardToInt(board)] = 0
	ops := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for i := 0; i < len(queue); i++ {
		cur := queue[i]
		times := visited[boardToInt(cur)]
		for j := 0; j < n; j++ {
			for k := 0; k < m; k++ {
				if cur[j][k] != 0 {
					continue
				}
				for _, op := range ops {
					row := j + op[0]
					col := k + op[1]
					if row < 0 || row >= n || col < 0 || col >= m {
						continue
					}
					clone := make([][]int, n)
					for h := 0; h < n; h++ {
						clone[h] = make([]int, m)
						copy(clone[h], cur[h])
					}
					clone[j][k], clone[row][col] = clone[row][col], clone[j][k]
					//fmt.Println(clone, cur)
					tempInt := boardToInt(clone)
					if _, ok := visited[tempInt]; ok {
						continue
					}
					if tempInt == 123450 {
						return times + 1
					}
					queue = append(queue, clone)
					visited[tempInt] = times + 1
				}
			}
		}
	}

	return -1
}

func boardToInt(board [][]int) int {
	result := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			result = result*10 + board[i][j]
		}
	}

	return result
}
