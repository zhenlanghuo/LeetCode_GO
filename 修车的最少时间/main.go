package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(repairCars([]int{4, 2, 3, 1}, 12))
	fmt.Println(repairCars([]int{5, 1, 8}, 6))
	fmt.Println(repairCars([]int{1, 2}, 6))
}

func repairCars(ranks []int, cars int) int64 {

	check := func(mid int64) bool {
		sum := 0
		for _, rank := range ranks {
			sum += int(math.Sqrt(float64(int64(mid / int64(rank)))))
		}

		if sum >= cars {
			return true
		}

		return false
	}

	left, right := int64(0), int64(1e14)
	for left <= right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
