package main

import "fmt"

func main() {
	fmt.Println(minimumVisitedCells([][]int{{3, 4, 2, 1}, {4, 2, 3, 1}, {2, 1, 0, 0}, {2, 4, 0, 0}}))
	fmt.Println(minimumVisitedCells([][]int{{3, 4, 2, 1}, {4, 2, 1, 1}, {2, 1, 1, 0}, {3, 4, 1, 0}}))
	fmt.Println(minimumVisitedCells([][]int{{2, 1, 0}, {1, 0, 0}}))
}

type UF struct {
	fa []int
}

func NewUF(size int) *UF {
	fa := make([]int, size)
	for i := 0; i < size; i++ {
		fa[i] = i
	}
	return &UF{fa: fa}
}

func (u *UF) Find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.Find(u.fa[x])
	}
	return u.fa[x]
}

func (u *UF) Merge(from, to int) {
	u.fa[from] = to
}

func minimumVisitedCells(grid [][]int) int {

	rowLen, colLen := len(grid), len(grid[0])
	ans := 0
	type pair struct{ i, j int }
	q := []pair{{0, 0}}
	rowUfs := make([]*UF, rowLen)
	for i := 0; i < rowLen; i++ {
		rowUfs[i] = NewUF(colLen + 1)
	}
	colUfs := make([]*UF, colLen)
	for j := 0; j < colLen; j++ {
		colUfs[j] = NewUF(rowLen + 1)
	}
	for len(q) != 0 {
		temp := q
		q = make([]pair, 0)
		for _, p := range temp {
			i, j := p.i, p.j
			//fmt.Println("i,j", i, j)
			g := grid[i][j]
			if i == rowLen-1 && j == colLen-1 {
				return ans + 1
			}

			// row
			next := rowUfs[i].Find(j + 1)
			for next < colLen && next <= j+g {
				q = append(q, pair{i, next})
				rowUfs[i].Merge(next, next+1)
				next = next + 1
			}
			//fmt.Println(rowUfs[i], q)

			// col
			next = colUfs[j].Find(i + 1)
			for next < rowLen && next <= i+g {
				q = append(q, pair{next, j})
				colUfs[j].Merge(next, next+1)
				next = next + 1
			}
			//fmt.Println(colUfs[j])
		}
		ans++
	}

	return -1
}
