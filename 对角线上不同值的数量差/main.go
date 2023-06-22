package main

func main() {

}

func differenceOfDistinctValues(grid [][]int) [][]int {
	n, m := len(grid), len(grid[0])
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			leftTop := make(map[int]struct{})
			leftTopX, leftTopY := i-1, j-1
			for leftTopX >= 0 && leftTopY >= 0 {
				leftTop[grid[leftTopX][leftTopY]] = struct{}{}
				leftTopX--
				leftTopY--
			}
			rightBottom := make(map[int]struct{})
			rightBottomX, rightBottomY := i+1, j+1
			for rightBottomX < n && rightBottomY < m {
				rightBottom[grid[rightBottomX][rightBottomY]] = struct{}{}
				rightBottomX++
				rightBottomY++
			}
			ans[i][j] = abs(len(leftTop) - len(rightBottom))
		}
	}

	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
