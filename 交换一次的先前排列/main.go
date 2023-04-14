package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(prevPermOpt1([]int{1, 9, 4, 6, 7}))
	fmt.Println(prevPermOpt1([]int{3, 2, 1}))
	fmt.Println(prevPermOpt1([]int{1, 1, 5}))
	fmt.Println(prevPermOpt1([]int{1, 7, 5, 4, 4, 6, 8}))
}

func prevPermOpt1(arr []int) []int {

	n := len(arr)
	right := len(arr) - 1
	for right-1 >= 0 && arr[right] >= arr[right-1] {
		right--
	}

	if right == 0 {
		return arr
	}

	pivot := sort.SearchInts(arr[right:n], arr[right-1])

	pivot = pivot + right - 1

	for pivot-1 >= 0 && arr[pivot] == arr[pivot-1] {
		pivot--
	}

	arr[right-1], arr[pivot] = arr[pivot], arr[right-1]

	return arr
}
