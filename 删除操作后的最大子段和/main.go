package main

import "fmt"

func main() {
	fmt.Println(maximumSegmentSum([]int{1, 2, 5, 6, 1}, []int{0, 3, 2, 4, 1}))
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

func maximumSegmentSum(nums []int, removeQueries []int) []int64 {
	n := len(nums)
	ans := make([]int64, n)
	sum := make([]int64, n+1)
	uf := NewUF(n + 1)
	for i := n - 1; i > 0; i-- {
		x := removeQueries[i]
		to := uf.Find(x + 1)
		uf.Merge(x, to)
		sum[to] += sum[x] + int64(nums[x])
		ans[i-1] = max(ans[i], sum[to])
	}

	return ans
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
