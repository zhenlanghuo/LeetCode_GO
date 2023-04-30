package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fieldOfGreatestBlessing([][]int{{0, 0, 1}, {1, 0, 1}}))
}

func fieldOfGreatestBlessing(forceField [][]int) int {

	xm := make(map[int]bool)
	xSlice := make([]int, 0)
	ym := make(map[int]bool)
	ySlice := make([]int, 0)
	for _, force := range forceField {
		x, y, size := force[0], force[1], force[2]
		if !xm[x*2-size] {
			xm[x*2-size] = true
			xSlice = append(xSlice, x*2-size)
		}
		if !xm[x*2+size] {
			xm[x*2+size] = true
			xSlice = append(xSlice, x*2+size)
		}
		if !ym[y*2-size] {
			ym[y*2-size] = true
			ySlice = append(ySlice, y*2-size)
		}
		if !ym[y*2+size] {
			ym[y*2+size] = true
			ySlice = append(ySlice, y*2+size)
		}
	}

	sort.Ints(xSlice)
	sort.Ints(ySlice)

	fmt.Println(xSlice)
	fmt.Println(ySlice)

	n, m := len(xSlice), len(ySlice)
	diff := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		diff[i] = make([]int, m+2)
	}
	for _, force := range forceField {
		x, y, size := force[0], force[1], force[2]
		left := sort.SearchInts(xSlice, x*2-size)
		right := sort.SearchInts(xSlice, x*2+size)
		top := sort.SearchInts(ySlice, y*2-size)
		bottom := sort.SearchInts(ySlice, y*2+size)
		diff[left+1][top+1] += 1     // left top
		diff[left+1][bottom+2] -= 1  // left bottom+1
		diff[right+2][bottom+2] += 1 // right+1 bottom+1
		diff[right+2][top+1] -= 1    // right+1 top
	}
	fmt.Println(diff)

	ans := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			diff[i][j] += diff[i-1][j] + diff[i][j-1] - diff[i-1][j-1]
			if diff[i][j] > ans {
				ans = diff[i][j]
			}
		}
	}

	return ans
}
