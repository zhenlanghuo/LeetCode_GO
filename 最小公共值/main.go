package main

func main() {

}

func getCommon(nums1 []int, nums2 []int) int {

	index1, index2 := 0, 0
	for index1 < len(nums1) && index2 < len(nums2) {
		if nums1[index1] == nums2[index2] {
			return nums1[index1]
		}

		if nums1[index1] < nums2[index2] {
			index1++
		} else {
			index2++
		}
	}

	return -1
}
