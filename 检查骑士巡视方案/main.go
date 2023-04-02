package main

func main() {

}

func checkValidGrid(grid [][]int) bool {

	n := len(grid)
	m := make([][]int, n*n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m[grid[i][j]] = []int{i, j}
		}
	}

	directions := [][]int{{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}

	if m[0][0] != 0 || m[0][1] != 0 {
		return false
	}

	for i := 0; i < n*n-1; i++ {
		x, y := m[i][0], m[i][1]
		_x, _y := m[i+1][0], m[i+1][1]

		flag := false
		for _, direct := range directions {
			dx, dy := direct[0], direct[1]
			if x+dx == _x && y+dy == _y {
				flag = true
			}
		}
		if !flag {
			return false
		}
	}

	return true
}
