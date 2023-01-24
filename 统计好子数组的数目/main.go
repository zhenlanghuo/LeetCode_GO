package main

import "fmt"

func main() {
	//fmt.Println(countGood([]int{1, 1, 1, 1, 1}, 10))
	//fmt.Println(countGood([]int{3, 1, 4, 3, 2, 2, 4}, 2))
	fmt.Println(countGood([]int{2, 1, 3, 1, 2, 2, 3, 3, 2, 2, 1, 1, 1, 3, 1}, 11))
}

//type static struct {
//	count int // 相同数字的数目
//	k     int // 相同数字的组合数目
//	index int // 在最大堆中的位置
//}

func countGood(nums []int, k int) int64 {
	staticMap := make(map[int]int)
	count := 0

	result := int64(0)
	right := 0
	staticMap[nums[0]] = 1
	for i := 0; i < len(nums); i++ {
		for right+1 < len(nums) {
			if count >= k {
				break
			}
			right = right + 1
			sta := staticMap[nums[right]]
			count -= (sta * (sta - 1)) / 2
			sta++
			count += (sta * (sta - 1)) / 2
			staticMap[nums[right]] = sta
		}

		//fmt.Println(i, right, count, staticMap)

		if count >= k {
			result += int64(len(nums) - right)
		}

		sta := staticMap[nums[i]]
		count -= (sta * (sta - 1)) / 2
		sta--
		count += (sta * (sta - 1)) / 2
		staticMap[nums[i]] = sta
		//fmt.Println(staticMap)
	}

	return result
}
