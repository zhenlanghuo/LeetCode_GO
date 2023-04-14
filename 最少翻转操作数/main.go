package main

import "fmt"

func main() {
	fmt.Println(minReverseOperations(4, 0, []int{1, 2}, 4))
}

type UF struct {
	fa map[int]int
}

func NewUF(start int, end int) *UF {
	fa := make(map[int]int)
	for i := start; i < end; i += 2 {
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

func minReverseOperations(n int, p int, banned []int, k int) []int {

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}
	bannedMap := make(map[int]bool)
	for _, ban := range banned {
		bannedMap[ban] = true
	}
	bannedMap[p] = true

	ufs := [2]*UF{NewUF(0, n+2), NewUF(1, n+2)}

	q := []int{p}
	for step := 0; len(q) != 0; step++ {
		temp := q
		q = nil
		for _, i := range temp {
			//fmt.Println(i)
			ans[i] = step
			mn := max(i-k+1, k-i-1)
			mx := min(i+k-1, n*2-k-i-1)
			uf := ufs[mn%2]
			//fmt.Println(uf, mn, mx)
			for j := uf.Find(mn); j <= mx; j = uf.Find(j + 2) {
				//fmt.Println(uf)
				if !bannedMap[j] {
					q = append(q, j)
				}
				uf.Merge(j, j+2) // 删除 j
			}
		}
		//fmt.Println(q)
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
