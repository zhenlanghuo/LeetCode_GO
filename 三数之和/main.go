package main

import (
	"fmt"
	"sort"
)

type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func main() {
	fmt.Println(threeSum([]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}))
}

func threeSum(nums []int) [][]int {
	sort.Sort(IntSlice(nums))

	result := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		for i !=0 && i < len(nums)-2 && nums[i] == nums[i-1] {
			i++
		}

		if nums[i] > 0 {
			return result
		}

		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				}
				for k--; j < k && nums[k] == nums[k+1]; k-- {
				}

				//for j < k && nums[j] == nums[j+1] {
				//	j++
				//}
				//for j < k && nums[k] == nums[k-1] {
				//	k--
				//}
				//j++
				//k--
			}

			if sum > 0 {
				k--
			}

			if sum < 0 {
				j++
			}
		}
	}

	return result
}
