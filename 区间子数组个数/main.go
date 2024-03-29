package main

import (
	"fmt"
)

func main() {
	fmt.Println(numSubarrayBoundedMax([]int{2, 1, 4, 3}, 2, 3))
	fmt.Println(numSubarrayBoundedMax([]int{2, 9, 2, 5, 6}, 2, 8))
	fmt.Println(numSubarrayBoundedMax([]int{876, 880, 482, 260, 132, 421, 744, 703, 795, 420, 871, 445, 400, 291, 358, 589, 617, 202, 755, 810, 227, 813, 549, 791, 418, 528, 835, 401, 526, 584, 873, 662, 13, 314, 988, 101, 299, 816, 833, 224, 160, 852, 179, 769, 646, 558, 661, 808, 651, 982, 878, 918, 406, 551, 467, 87, 139, 387, 16, 531, 307, 389, 939, 551, 613, 36, 528, 460, 404, 314, 66, 111, 458, 531, 944, 461, 951, 419, 82, 896, 467, 353, 704, 905, 705, 760, 61, 422, 395, 298, 127, 516, 153, 299, 801, 341, 668, 598, 98, 241},
		658, 719))
}

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	ans := 0

	//curMax := nums[0]
	prev := -1 // 最近大于right的下标
	prevCount := 0
	for i, num := range nums {
		if num > right {
			// 当 num 大于 right 时, num 不可能成为满足要求的子数组中的一员, 因此更新 prev
			prev = i
			prevCount = 0
		} else {
			if num >= left && num <= right {
				// 如果当前num在区间内，那么以num结束的子数组有 i-prev 个
				prevCount = i - prev
				ans += prevCount
			} else {
				// 如果当前num不在区间内，那么以num结束的子数组有 prevCount 个
				ans += prevCount
			}
		}
	}

	return ans
}

//func numSubarrayBoundedMax(nums []int, left int, right int) int {
//	n := len(nums)
//	i, j := 0, 0
//	tempMax := nums[0]
//	result := 0
//	for j < n {
//		tempMax = nums[j]
//		for j < n && tempMax <= right {
//			j++
//			if j < n {
//				tempMax = max(tempMax, nums[j])
//			}
//		}
//
//		//fmt.Println(i, j)
//		count := 0
//		for i < j {
//			if nums[i] >= left && nums[i] <= right {
//				result += (count + 1) * (j - i)
//				count = 0
//			} else {
//				count++
//			}
//			i++
//		}
//		j++
//		i = j
//	}
//
//	return result
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
