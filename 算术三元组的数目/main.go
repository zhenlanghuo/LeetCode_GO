package main

import "fmt"

func main() {
	fmt.Println(arithmeticTriplets([]int{0, 1, 4, 6, 7, 10}, 3))
	fmt.Println(arithmeticTriplets([]int{4, 5, 6, 7, 8, 9}, 2))
}

func arithmeticTriplets(nums []int, diff int) int {

	ans := 0

	m := make(map[int]int)
	for i, num := range nums {
		m[num] = i
	}
	fmt.Println(m)

	for i := 0; i < len(nums); i++ {
		if nums[i] == -1 {
			continue
		}

		count := 1
		for {
			v, ok := m[nums[i]+diff*count]
			if !ok {
				break
			}
			count++
			nums[v] = -1
		}
		//fmt.Println(i, nums, count)

		if count >= 3 {
			ans += count - 2
		}
	}

	return ans
}
