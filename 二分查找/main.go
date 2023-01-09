package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	for i := 0; i < 1e4; i++ {
		array := make([]int, 100)
		for j := 0; j < len(array); j++ {
			array[j] = rand.Intn(10) + 1
		}
		sort.Ints(array)
		target := rand.Intn(15)

		index1 := findLastIndexEGTV1(array, target)
		index2 := findLastIndexEGTV2(array, target)

		if index1 == index2 {
			fmt.Println(target, array)
			fmt.Println(index1, index2)
		}
	}

	fmt.Println("finish")

	//array := []int{2, 3, 3, 4, 5, 6, 7, 8, 9, 10, 11, 11}
	//fmt.Println(findLastIndexEGTV2(array, 3))

}

func findFirstIndexEGTV1(nums []int, x int) int {
	left, right := 0, len(nums)
	// 搜索区间为[left, right), left=right时结束
	for left < right {
		mid := (left + right) / 2
		if nums[mid] >= x {
			// nums[mid]>x时, 目标只能在区间[left, mid)中
			// nums[mid]=x时, 目标只能在区间[left, mid)中
			right = mid
		} else if nums[mid] < x {
			// nums[mid] < x, 目标只能在区间[mid+1, right)中
			left = mid + 1
		}
	}

	// 碰到nums[mid]=x时, right=mid, 而搜索的右区间是开区间, left=right时结束, 所以如果mid为答案, 最终结束时为left=right=mid
	return left
}

func findFirstIndexEGTV2(nums []int, x int) int {
	left, right := 0, len(nums)-1
	// 搜索区间为[left, right, left>right时结束
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= x {
			// nums[mid]>x时, 目标只能在区间[left, mid-1]中
			// nums[mid]=x时, 目标只能在区间[left, mid-1]中
			right = mid - 1
		} else if nums[mid] < x {
			// nums[mid] < x, 目标只能在区间[mid+1, right]中
			left = mid + 1
		}
	}

	// nums[mid]=x时, right = mid - 1, 搜索的右区间是闭区间, left > right时结束, 所以如果mid为答案, 最终结束时为left=mid>right=mid-1)
	return left
}

func findLastIndexEGTV1(nums []int, x int) int {
	left, right := 0, len(nums)
	// 搜索区间为[left, right), left=right时结束
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > x {
			// nums[mid]>x时, 目标只能在区间[left, mid)中
			right = mid
		} else if nums[mid] <= x {
			// nums[mid]<x, 目标只能在区间[mid+1, right)中
			// nums[mid]=x时, 目标只能在区间[mid+1, right)中
			left = mid + 1
		}
	}

	// nums[mid]=x时, left = mid + 1, 搜索的右区间是开区间, left = right 时结束, 所以如果mid为答案, 最终结束时为left=mid+1=right
	return left - 1
}

func findLastIndexEGTV2(nums []int, x int) int {
	left, right := 0, len(nums)-1
	// 搜索区间为[left, right], left>right时结束
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > x {
			// nums[mid]>x时, 目标只能在区间[left, mid-1]中
			right = mid - 1
		} else if nums[mid] <= x {
			// nums[mid]<x, 目标只能在区间[mid+1, right]中
			// nums[mid]=x时, 目标只能在区间[mid+1, right]中
			left = mid + 1
		}
	}

	// nums[mid]=时, left = mid + 1, 搜索的右区间是闭区间, left > right 时结束, 所以如果mid为答案, 最终结束时 left=mid+1>right=mid
	return left - 1
}
