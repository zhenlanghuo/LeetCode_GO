package main

import "fmt"

func main() {
	fmt.Println(subarrayLCM([]int{3, 6, 2, 7, 1}, 6))
	fmt.Println(subarrayLCM([]int{2, 3}, 6))
}

func subarrayLCM(nums []int, k int) int {

	ans := 0
	for i, _ := range nums {
		lcm := 1
		for j := i; j < len(nums); j++ {
			lcm = lcm * nums[j] / gcd(lcm, nums[j])
			//fmt.Println(i, j, lcm)
			if k%lcm > 0 {
				break
			}
			if lcm == k {
				ans++
			}
		}
	}

	return ans
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
