package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3,4}))
}

var (
	result [][]int
)

func permute(nums []int) [][]int {
	result = make([][]int, 0)
	permute_(nums, 0)
	return result
}

//func permute_(nums []int, level int, result *[][]int) {
//	if level == len(nums) {
//		temp := make([]int, 0, len(nums))
//		for i:=0;i<len(nums);i++{
//			temp = append(temp, nums[i])
//		}
//		*result = append(*result, temp)
//		//fmt.Println(nums)
//	}
//
//	for i := level; i < len(nums); i++ {
//		nums[level], nums[i] = nums[i], nums[level]
//		permute_(nums, level+1, result)
//		nums[level], nums[i] = nums[i], nums[level]
//	}
//}

func permute_(nums []int, level int) {
	//fmt.Println(nums, level)
	if level == len(nums) {
		//temp := make([]int, 0, len(nums))
		//for i:=0;i<len(nums);i++{
		//	temp = append(temp, nums[i])
		//}
		//*result = append(*result, temp)
		fmt.Println(nums)
		dst := make([]int, len(nums))
		copy(dst, nums)
		result = append(result, dst)
	}

	for i := level; i < len(nums); i++ {
		if i != level && nums[level] == nums[i] {
			continue
		}
		nums[level], nums[i] = nums[i], nums[level]
		permute_(nums, level+1)
		nums[level], nums[i] = nums[i], nums[level]
	}
}
