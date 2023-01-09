package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 4, 5, 6, 0}
	nums2 := []int{3}
	merge(nums1, 5, nums2, 1)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := n + m - 1; i > n-1; i-- {
		nums1[i] = nums1[i-n]
	}

	fmt.Println(nums1)

	count := 0
	i, j := n, 0
	for count != m+n {
		if i >= m+n {
			nums1[count] = nums2[j]
			j++
		} else if j >= n {
			nums1[count] = nums1[i]
			i++
		} else {
			if nums1[i] < nums2[j] {
				nums1[count] = nums1[i]
				i++
			} else {
				nums1[count] = nums2[j]
				j++
			}
		}
		count++
	}

}
