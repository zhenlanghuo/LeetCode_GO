package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minInterval([][]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}}, []int{2, 3, 4, 5}))
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

type QueryPair struct {
	query int
	index int
}

func minInterval(intervals [][]int, queries []int) []int {

	m := len(queries)
	ans := make([]int, m)
	qp := make([]QueryPair, m)
	for i := 0; i < m; i++ {
		ans[i] = -1
		qp[i] = QueryPair{
			query: queries[i],
			index: i,
		}
	}

	sort.Slice(qp, func(i, j int) bool {
		return qp[i].query < qp[j].query
	})

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1]-intervals[i][0] < intervals[j][1]-intervals[j][0]
	})

	//fmt.Println(intervals)
	//fmt.Println(qp)

	uf := NewUF(m + 1)
	for _, interval := range intervals {
		l, r := interval[0], interval[1]
		start := sort.Search(m, func(i int) bool {
			return qp[i].query >= l
		})
		//fmt.Println(start, uf)
		for j := uf.Find(start); j < m && qp[j].query <= r; j = uf.Find(j + 1) {
			ans[qp[j].index] = r - l + 1
			uf.Merge(j, j+1)
		}
	}

	return ans
}
