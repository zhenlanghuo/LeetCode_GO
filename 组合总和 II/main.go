package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{2,5,2,1,2}, 5))
}

var (
	result [][]int
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result = make([][]int, 0)
	solve(candidates, 0, make([]int, 0), 0, target)
	return result
}

func solve(nums []int, start int, temp []int, tempSum, target int) {
	if start == len(nums) {
		//fmt.Println(temp)
		if target == tempSum {
			clone := make([]int, len(temp))
			copy(clone, temp)
			result = append(result, clone)
		}
		return
	}

	end := len(nums)
	for i := start + 1; i < len(nums); i++ {
		if nums[i] != nums[start] {
			end = i
			break
		}
	}

	i := start
	for ; i < end; i++ {
		if tempSum+nums[start] > target {
			break
		}
		temp = append(temp, nums[start])
		tempSum += nums[start]
		solve(nums, end, temp, tempSum, target)
	}
	temp = temp[:len(temp)-(i-start)]
	//fmt.Println(start)
	tempSum -= nums[start] * (i - start)
	solve(nums, end, temp, tempSum, target)
}
