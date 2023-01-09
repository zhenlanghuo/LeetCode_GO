package main

func main() {

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen%2 != 0 {
		k := totalLen / 2
		return float64(findKthElement(nums1, nums2, k+1))
	} else {
		k := totalLen / 2
		return (float64(findKthElement(nums1, nums2, k)) + float64(findKthElement(nums1, nums2, k+1))) / float64(2)
	}
}

func findKthElement(nums1 []int, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}

		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}

		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}

		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		if nums1[newIndex1] <= nums2[newIndex2] {
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
