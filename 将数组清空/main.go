package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(countOperationsToEmptyArray([]int{3, 4, -1}))
	fmt.Println(countOperationsToEmptyArray([]int{1, 2, 3}))
	fmt.Println(countOperationsToEmptyArray([]int{1, 2, 4, 3}))
	fmt.Println(countOperationsToEmptyArray([]int{4, 3, 1, 2, 5}))
	fmt.Println(countOperationsToEmptyArray([]int{6, 18, 13, -15}))
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

func countOperationsToEmptyArray(nums []int) int64 {

	ans := int64(0)
	n := len(nums)
	clone := make([]int, n)
	copy(clone, nums)

	sort.Ints(clone)

	uf := NewUF(n)
	start := 0

	for i := 0; i < n-1; i++ {
		mn := clone[i]
		for mn != nums[start] {
			start = uf.Find((start + 1) % n)
			//end = uf.Find((end + 1) % n)
			ans++
		}
		//fmt.Println(clone[i], uf.fa, ans, start, end)
		ans++
		uf.Merge(start, uf.Find((start+1)%n))
		start = uf.Find((start + 1) % n)
		//fmt.Println(clone[i], uf.fa, ans, start, end)
	}
	ans++

	return ans
}
