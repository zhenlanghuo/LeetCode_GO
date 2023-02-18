package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(advantageCount([]int{2, 7, 11, 15}, []int{1, 10, 4, 11}))
	fmt.Println(advantageCount([]int{12, 24, 8, 32}, []int{13, 25, 32, 11}))
}

func advantageCount(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	ans := make([]int, n)
	nums1Index := make([]int, n)
	nums2Index := make([]int, n)
	for i := 0; i < n; i++ {
		nums1Index[i] = i
		nums2Index[i] = i
	}

	sort.Slice(nums1Index, func(i, j int) bool {
		return nums1[nums1Index[i]] < nums1[nums1Index[j]]
	})

	sort.Slice(nums2Index, func(i, j int) bool {
		return nums2[nums2Index[i]] < nums2[nums2Index[j]]
	})

	//for i := 0; i < n; i++ {
	//	fmt.Printf("%v,", nums1[nums1Index[i]])
	//}
	//fmt.Println(nums1Index)
	//
	//for i := 0; i < n; i++ {
	//	fmt.Printf("%v,", nums2[nums2Index[i]])
	//}
	//fmt.Println(nums2Index)

	l, r, i := 0, n-1, 0
	for i < n {
		if nums2[nums2Index[l]] < nums1[nums1Index[i]] {
			ans[nums2Index[l]] = nums1[nums1Index[i]]
			l++
		} else {
			ans[nums2Index[r]] = nums1[nums1Index[i]]
			r--
		}
		i++
	}

	return ans
}
