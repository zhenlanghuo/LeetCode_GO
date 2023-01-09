package main

import (
	"fmt"
)

func main() {
	arr := []int{2147483647, 2147483647, -2147483647, -2147483647, -2147483647, 2147483647}
	fmt.Println(reversePairs(arr))
	fmt.Println(arr)
}

var temp []int
var count int

func reversePairs(nums []int) int {
	temp = make([]int, len(nums))
	count = 0
	mergeSort(nums, 0, len(nums)-1)

	return count
}

func mergeSort(nums []int, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := (lo + hi) / 2
	mergeSort(nums, lo, mid)
	mergeSort(nums, mid+1, hi)
	merge(nums, lo, hi, mid)
}

func merge(nums []int, lo, hi, mid int) {
	for i := lo; i <= hi; i++ {
		temp[i] = nums[i]
	}

	j := mid + 1
	for i := lo; i <= mid; i++ {
		for j <= hi && temp[i] > 2*temp[j] {
			j++
		}
		count += j - mid - 1
	}

	l, r := lo, mid+1
	for lo <= hi {
		if l > mid {
			nums[lo] = temp[r]
			r++
		} else if r > hi {
			nums[lo] = temp[l]
			//count += sort.SearchInts(temp[mid+1:r], (temp[l]+1)/2)
			l++
		} else {
			if temp[l] >= temp[r] {
				nums[lo] = temp[r]
				r++
			} else {
				nums[lo] = temp[l]
				//count += sort.SearchInts(temp[mid+1:r+1], (temp[l]+1)/2)
				l++
			}
		}
		lo++
	}
}
