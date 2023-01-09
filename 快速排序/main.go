package main

import (
	"fmt"
)

func main() {
	list := []int{1, 38, 65, 44, 4, 8, 33, 1, 22, -11, 6, 34, 55, 54, 9}
	quickSort(list)
	fmt.Println(list)

	//count := 0
	//for i := 0; i < 10000; i++ {
	//	array1 := make([]int, 100)
	//	array2 := make([]int, 100)
	//	for j := 0; j < 100; j++ {
	//		ran := rand.Intn(1000)
	//		array1[j] = ran
	//		array2[j] = ran
	//	}
	//
	//	//fmt.Println(array1)
	//	//fmt.Println(array2)
	//	//fmt.Println()
	//
	//	quickSort(array1)
	//	sort.Ints(array2)
	//	if !equal(array1, array2) {
	//		count++
	//	}
	//}
	//
	//fmt.Println(count)
}

func equal(array1, array2 []int) bool {
	for i := 0; i < len(array1); i++ {
		if array1[i] != array2[i] {
			return false
		}
	}

	return true
}

func partition(arr []int, startIndex, endIndex int) int {
	var (
		pivot = arr[startIndex]
		left  = startIndex
		right = endIndex
	)

	for left != right {
		for left < right && pivot < arr[right] {
			right--
		}

		for left < right && pivot >= arr[left] {
			left++
		}

		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}

	arr[startIndex], arr[left] = arr[left], arr[startIndex]

	return left
}

//func partition(arr []int, startIndex, endIndex int) int {
//	var (
//		pivot = arr[startIndex]
//		left  = startIndex
//		right = endIndex
//	)
//
//	for left != right {
//		for left < right && arr[left] <= pivot {
//			left++
//		}
//
//		for left < right && arr[right] > pivot {
//			right--
//		}
//
//		if left < right {
//			arr[left], arr[right] = arr[right], arr[left]
//		}
//	}
//
//	if arr[left] > pivot {
//		left = left - 1
//	}
//
//	arr[startIndex], arr[left] = arr[left], arr[startIndex]
//	return left
//}

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivotIndex := partition(arr, 0, len(arr)-1)

	fmt.Println(arr, arr[0:pivotIndex], arr[pivotIndex+1:])

	quickSort(arr[0:pivotIndex])
	quickSort(arr[pivotIndex+1:])
}
