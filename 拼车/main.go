package main

import "fmt"

func main() {
}

func carPooling(trips [][]int, capacity int) bool {
	n := 0

	for i := 0; i < len(trips); i++ {
		if n < trips[i][2] {
			n = trips[i][2] + 1
		}
	}

	diff := make([]int, n+1)
	count := make([]int, n+1)

	for i := 0; i < len(trips); i++ {
		add(diff, trips[i][1], trips[i][2]-1, trips[i][0])
	}

	fmt.Println(diff)

	for i := 0; i < len(count); i++ {
		if i == 0 {
			count[i] = diff[i]
		} else {
			count[i] = count[i-1] + diff[i]
		}
		if count[i] > capacity {
			return false
		}
	}

	return true
}

func add(diff []int, i, j, v int) {
	diff[i] += v
	if j+1 < len(diff) {
		diff[j+1] -= v
	}
}
