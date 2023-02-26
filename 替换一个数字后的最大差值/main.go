package main

import "fmt"

func main() {
	fmt.Println(minMaxDifference(11891))
}

func minMaxDifference(num int) int {
	maxNum, minNum := num, num
	nums := make([]int, 0)
	numsMap := make(map[int]bool)
	for num > 0 {
		nums = append(nums, num%10)
		numsMap[num%10] = true
		num = num / 10
	}

	reverse(nums)

	fmt.Println(nums)
	fmt.Println(numsMap)

	for i := 0; i <= 9; i++ {
		for k, _ := range numsMap {
			temp := 0
			for _, v := range nums {
				if v == k {
					v = i
				}
				temp = temp*10 + v
			}
			//fmt.Println(i, k, temp)
			maxNum = max(maxNum, temp)
			minNum = min(minNum, temp)
		}
	}

	fmt.Println(maxNum, minNum)

	return maxNum - minNum
}

func reverse(arr []int) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
