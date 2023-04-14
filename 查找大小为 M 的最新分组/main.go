package main

import "fmt"

func main() {
	fmt.Println(findLatestStep([]int{3, 5, 1, 2, 4}, 1))
	fmt.Println(findLatestStep([]int{3, 1, 5, 4, 2}, 2))
	fmt.Println(findLatestStep([]int{1}, 1))
	fmt.Println(findLatestStep([]int{2, 1}, 2))
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

func findLatestStep(arr []int, m int) int {

	n := len(arr)
	count := make([]int, n+1)
	countMap := make(map[int]int)
	//oneArr := make([]int, n)
	ans := -1

	uf := NewUF(n + 1)
	for k, i := range arr {
		i = i - 1
		if countMap[count[i]] > 0 {
			countMap[count[i]]--
		}
		count[i]++
		countMap[count[i]]++
		to := uf.Find(i + 1)
		uf.Merge(i, i+1)
		if countMap[count[to]] > 0 {
			countMap[count[to]]--
		}
		if countMap[count[i]] > 0 {
			countMap[count[i]]--
		}
		count[to] = count[to] + count[i]
		countMap[count[to]]++

		//fmt.Println(countMap, count)

		if countMap[m] > 0 {
			ans = k + 1
		}

	}

	return ans
}
