package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(splitNum(4325))
	fmt.Println(splitNum(687))

}

func splitNum(num int) int {
	nums := make([]int, 0)
	for num != 0 {
		nums = append(nums, num%10)
		num = num / 10
	}

	fmt.Println(nums)
	sort.Ints(nums)

	num1 := nums[0]
	num2 := nums[1]

	for i := 2; i < len(nums); i++ {
		fmt.Println(num1, num2)
		if num1 < num2 {
			num1 = num1*10 + nums[i]
		} else {
			num2 = num2*10 + nums[i]
		}
	}

	return num1 + num2
}
