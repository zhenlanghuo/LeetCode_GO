package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(makeSubKSumEqual([]int{1, 4, 1, 3, 3, 4, 3, 4, 3}, 3))
	fmt.Println(makeSubKSumEqual([]int{1, 4, 1, 3}, 2))
	fmt.Println(makeSubKSumEqual([]int{2, 5, 5, 7}, 3))
	//fmt.Println(makeSubKSumEqual([]int{9}, 1))
}

func makeSubKSumEqual(arr []int, k int) int64 {

	n := len(arr)
	m := make(map[int]int)
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		m[i] = (i + k) % n
	}

	group := make([][]int, 0)
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}

		cur := i
		temp := []int{cur}
		visited[cur] = true
		for !visited[m[cur]] {
			temp = append(temp, m[cur])
			visited[m[cur]] = true
			cur = m[cur]
		}

		group = append(group, temp)
	}

	ans := int64(0)

	fmt.Println(group)

	for _, g := range group {
		temp := make([]int, len(g))
		for i := 0; i < len(g); i++ {
			temp[i] = arr[g[i]]
		}

		sort.Ints(temp)

		if len(temp)%2 == 1 {
			ans += diff(temp, temp[len(temp)/2])
		} else {
			ans += min(diff(temp, temp[len(temp)/2]), diff(temp, temp[(len(temp)/2)-1]))
		}
	}

	return ans
}

func diff(temp []int, num int) int64 {
	ans := int64(0)
	for _, t := range temp {
		ans += int64(abs(t - num))
	}
	return ans
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
