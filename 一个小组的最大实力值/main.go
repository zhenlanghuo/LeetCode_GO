package main

import (
	"fmt"
)

func main() {
	//fmt.Println(maxStrength([]int{3, -1, -5, 2, 5, -9}))
	fmt.Println(maxStrength([]int{0, 0, -9, 2}))
	fmt.Println(maxStrength([]int{0}))
}

func maxStrength(nums []int) int64 {
	ans := int64(1)
	//sort.Ints(nums)

	newNums := make([]int, 0, len(nums))
	zeroCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCount++
			continue
		}
		newNums = append(newNums, nums[i])
	}

	if len(newNums) == 0 {
		return 0
	}

	nCount := 0
	mx := -10
	for _, num := range newNums {
		if num < 0 {
			nCount++
			if num > mx {
				mx = num
			}
		}
		if num == 0 {
			continue
		}
		ans = ans * int64(num)
	}

	if len(newNums) == 1 && nCount%2 == 1 {
		if zeroCount != 0 {
			return 0
		}
		return ans
	}

	if nCount%2 == 1 {
		ans = ans / int64(mx)
	}

	return ans
}
