package main

import "fmt"

func main() {
	nums := []int{1,3,2}
	nextPermutation(nums)
	fmt.Println(nums)
}

func nextPermutation(nums []int) {
	n := len(nums)
	// 较小数, 要尽量靠右
	left := 0
	for i := n - 1; i >= 0; i-- {
		if nums[i] > nums[(i-1+n)%n] {
			left = (i - 1 + n) % n
			break
		}
	}
	fmt.Println(left)

	if left == n-1 {
		reverse(nums, 0, n-1)
		return
	}

	// 较大数，要尽可能小
	right := 0
	for i := n - 1; i >= 0; i-- {
		if nums[i] > nums[left] {
			right = i
			break
		}
	}

	//fmt.Println(right)
	nums[left], nums[right] = nums[right], nums[left]
	//fmt.Println(nums)
	reverse(nums, left+1, n-1)

	return
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

//func permutation(nums []int, level int) {
//	if level == len(nums) {
//		fmt.Println(nums)
//	}
//
//	for i := level; i < len(nums); i++ {
//		swap(nums, level, i)
//		//if level+ 1 < len(nums) && level != i {
//		//	swap(nums, level+1, i)
//		//}
//		permutation(nums, level+1)
//		//if level+ 1 < len(nums) && level != i {
//		//	swap(nums, level+1, i)
//		//}
//		swap(nums, level, i)
//	}
//}
//
//func swap(nums []int, i, j int) {
//	temp := nums[i]
//	nums[i] = nums[j]
//	nums[j] = temp
//}
