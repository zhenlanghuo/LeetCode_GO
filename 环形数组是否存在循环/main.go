package main

import (
	"fmt"
)

func main() {
	fmt.Println(circularArrayLoop([]int{1, -1, 5, 1, 4}))
}

func circularArrayLoop(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		if circularArrayLoop_(nums, i) {
			return true
		}
		fmt.Println(nums)
	}

	return false
}

func circularArrayLoop_(nums []int, i int) bool {
	defer func() {
		for nums[i] != 0 {
			nextIndex := next(nums, i)
			nums[i] = 0
			i = nextIndex
		}
	}()

	slow, fast := i, i
	for {
		slow = next(nums, slow)
		fast = next(nums, next(nums, fast))
		if slow == fast {
			break
		}
	}

	//slow = 0
	//for {
	//	slow = next(nums, slow)
	//	fast = next(nums, fast)
	//	if slow == fast {
	//		break
	//	}
	//}

	//fmt.Println(slow, fast)

	count := 0
	for {
		count++
		fast = next(nums, fast)
		//fmt.Println(fast)

		if !(nums[fast] > 0 && nums[slow] > 0 || nums[fast] < 0 && nums[slow] < 0) {
			return false
		}

		if slow == fast && count == 1 {
			return false
		}

		if slow == fast {
			break
		}
	}

	return true
}

func next(nums []int, index int) int {
	length := len(nums)
	step := nums[index] % length
	nextIndex := (index + step) % length
	if nextIndex < 0 {
		nextIndex = (nextIndex + length) % length
	}

	return nextIndex
}
