package main

import "fmt"

func main() {

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

func countCompleteComponents(n int, edges [][]int) int {
	ans := 0

	uf := NewUF(n)
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
	}

	for _, edge := range edges {
		x, y := edge[0], edge[1]
		if x > y {
			x, y = y, x
		}
		uf.Merge(uf.Find(x), uf.Find(y))
		g[x][y] = 1
		g[y][x] = 1
	}

	fmt.Println(uf)
	fmt.Println(g)

	set := make(map[int][]int)
	for i := 0; i < n; i++ {
		set[uf.Find(i)] = append(set[uf.Find(i)], i)
	}

	fmt.Println(set)

	for _, nodes := range set {
		flag := true
		for i := 0; i < len(nodes); i++ {
			for j := i; j < len(nodes); j++ {
				if g[nodes[i]][nodes[j]] != 1 {
					flag = false
				}
			}
		}
		if flag {
			ans++
		}
	}

	return ans
}
