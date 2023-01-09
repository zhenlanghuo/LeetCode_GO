package main

import "fmt"

func main() {
	list := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(list, 4))
}

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	buildMaxHead(nums, n)
	fmt.Println(nums)
	for i := 1; i < k; i++ {
		nums[0], nums[n-i] = nums[n-i], nums[0]
		maxHeapify(nums, 0, n-i)
		fmt.Println(nums)
	}

	return nums[0]
}

//func findKthLargest(nums []int, k int) int {
//	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
//}

func quickSelect(nums []int, start, end, target int) int {
	q := partition(nums, start, end)
	if q == target {
		return nums[q]
	}

	if q < target {
		return quickSelect(nums, q+1, end, target)
	} else {
		return quickSelect(nums, start, q-1, target)
	}
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

func buildMaxHead(arr []int, headSize int) {
	for i := headSize / 2; i >= 0; i-- {
		maxHeapify(arr, i, headSize)
	}
}

func maxHeapify(arr []int, index, headSize int) {
	left, right, largest := index*2+1, index*2+2, index
	if left < headSize && arr[left] > arr[largest] {
		largest = left
	}
	if right < headSize && arr[right] > arr[largest] {
		largest = right
	}
	if largest != index {
		arr[index], arr[largest] = arr[largest], arr[index]
		maxHeapify(arr, largest, headSize)
	}
}
