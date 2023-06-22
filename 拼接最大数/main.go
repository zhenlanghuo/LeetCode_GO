package main

func main() {

}

func maxNumber(nums1, nums2 []int, k int) (res []int) {
	start := 0
	if k > len(nums2) {
		start = k - len(nums2)
	}
	for i := start; i <= k && i <= len(nums1); i++ {
		s1 := getMaxSequence(nums1, i)
		s2 := getMaxSequence(nums2, k-i)
		merged := merge(s1, s2)
		if lexicographicalLess(res, merged) {
			res = merged
		}
	}
	return
}

func lexicographicalLess(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}

func merge(a, b []int) []int {
	merged := make([]int, len(a)+len(b))
	for i := range merged {
		if lexicographicalLess(a, b) {
			merged[i], b = b[0], b[1:]
		} else {
			merged[i], a = a[0], a[1:]
		}
	}
	return merged
}

func getMaxSequence(nums []int, k int) []int {
	n := len(nums)
	stack := make([]int, 0, k)

	for i := 0; i < n; i++ {
		for len(stack) > 0 && len(stack)+n-i > k && nums[i] > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) < k {
			stack = append(stack, nums[i])
		}
	}

	return stack
}
