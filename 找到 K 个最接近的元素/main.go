package main

import (
	"fmt"
)

func main() {
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, -1))

	fmt.Println(searchFirstGETIndex([]int{1, 2, 2, 2, 4, 6, 9}, 3))
}

func findClosestElements(arr []int, k int, x int) []int {

	//sort.Ints(arr)
	insertIndex := searchFirstGETIndex(arr, x)
	fmt.Println(insertIndex)
	if insertIndex == len(arr) {
		return arr[len(arr)-k:]
	}

	leftIndex, rightIndex := 0, 0
	if arr[insertIndex] == x {
		leftIndex = insertIndex - 1
		rightIndex = insertIndex + 1
	} else {
		leftIndex = insertIndex - 1
		rightIndex = insertIndex
	}

	for rightIndex-leftIndex-1 < k {
		if rightIndex >= len(arr) {
			return arr[len(arr)-k:]
		}

		if leftIndex < 0 {
			return arr[:k]
		}

		if x-arr[leftIndex] == arr[rightIndex]-x {
			leftIndex--
		} else if x-arr[leftIndex] < arr[rightIndex]-x {
			leftIndex--
		} else {
			rightIndex++
		}
	}

	return arr[leftIndex+1 : rightIndex]
}

func searchFirstGETIndex(array []int, x int) int {
	start, end := 0, len(array)
	for start < end {
		mid := (start + end) / 2
		if array[mid] < x {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return start
}
