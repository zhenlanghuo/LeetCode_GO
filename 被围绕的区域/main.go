package main

func main() {

}

type UF struct {
	parent []int
	count  int
}

func newUF(n int) *UF {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UF{
		parent: parent,
		count:  n,
	}
}

func (u *UF) union(x, y int) {
	rootX := u.find(x)
	rootY := u.find(y)
	if rootX == rootY {
		return
	}

	u.parent[rootX] = rootY
	u.count--
}

func (u *UF) find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.find(u.parent[x])
	}
	return u.parent[x]
}

func (u *UF) connected(x, y int) bool {
	return u.find(x) == u.find(y)
}

func solve(board [][]byte) {
	row, col := len(board), len(board[0])

	uf := newUF(row*col + 1)
	dummy := row * col

	for i := 0; i < row; i++ {
		if board[i][0] == 'O' {
			uf.union(i*col, dummy)
		}
		if board[i][col-1] == 'O' {
			uf.union(i*col+col-1, dummy)
		}
	}

	for i := 0; i < col; i++ {
		if board[0][i] == 'O' {
			uf.union(i, dummy)
		}

		if board[row-1][i] == 'O' {
			uf.union((row-1)*col+i, dummy)
		}
	}

	drs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 1; i < row-1; i++ {
		for j := 1; j < col-1; j++ {
			if board[i][j] != 'O' {
				continue
			}

			for _, dr := range drs {
				nRow := i + dr[0]
				nCol := j + dr[1]

				if board[nRow][nCol] == 'O' {
					uf.union(i*col+j, nRow*col+nCol)
				}
			}
		}
	}

	for i := 1; i < row-1; i++ {
		for j := 1; j < col-1; j++ {
			if board[i][j] != 'O' {
				continue
			}

			if !uf.connected(i*col+j, dummy) {
				board[i][j] = 'X'
			}
		}
	}
}
