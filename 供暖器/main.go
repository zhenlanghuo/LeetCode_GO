package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findRadius([]int{1,5}, []int{2}))
}

func findRadius(houses []int, heaters []int) int {

	sort.Ints(houses)
	sort.Ints(heaters)

	max := 0

	j := 0
	for i := 0; i < len(houses); i++ {
		for j < len(heaters)-1 {
			cur := abs(houses[i] - heaters[j])
			next := abs(houses[i] - heaters[j+1])
			if cur < next {
				break
			}
			j++
		}

		if max < abs(houses[i]-heaters[j]) {
			max = abs(houses[i] - heaters[j])
		}
	}

	return max
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return a * -1
}
