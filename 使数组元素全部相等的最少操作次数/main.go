package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(minOperations([]int{3, 1, 6, 8}, []int{5, 1, 1, 5}))
	fmt.Println(minOperations([]int{2, 9, 6, 3}, []int{10, 1, 2, 3, 5, 11}))
	//fmt.Println(minOperations([]int{47, 50, 97, 58, 87, 72, 41, 63, 41, 51, 17, 21, 7, 100, 69, 66, 79, 92, 84, 9, 57, 26, 26, 28, 83, 38}, []int{50, 84, 76, 41, 64, 82, 20, 22, 64, 7, 38, 92, 39, 28, 22, 3, 41, 46, 47, 50, 88, 51, 9, 49, 38, 67, 26, 65, 89, 27, 71, 25, 77, 72, 65, 41, 84, 68, 51, 26, 84, 24, 79, 41, 96, 83, 92, 9, 93, 84, 35, 70, 74, 79, 37, 38, 26, 26, 41, 26}))
}

func minOperations(nums []int, queries []int) []int64 {

	sort.Ints(nums)
	//fmt.Println(nums)

	n, m := len(nums), len(queries)
	queriesIndex := make([]int, m)
	for i := 0; i < m; i++ {
		queriesIndex[i] = i
	}
	sort.Slice(queriesIndex, func(i, j int) bool {
		return queries[queriesIndex[i]] < queries[queriesIndex[j]]
	})

	//fmt.Println(queriesIndex)

	ans := make([]int64, m)
	left := -1
	for i := 0; i < n; i++ {
		ans[queriesIndex[0]] += int64(abs(nums[i] - queries[queriesIndex[0]]))
		if nums[i] <= queries[queriesIndex[0]] {
			left = i
		}
	}

	//fmt.Println(ans)

	right := left + 1
	preTemp := int64(0)
	for j := 1; j < m; j++ {
		//fmt.Println("@@@", queries[queriesIndex[j]], queries[queriesIndex[j-1]], left, right)

		temp := preTemp + int64((queries[queriesIndex[j]]-queries[queriesIndex[j-1]])*(right-left-1))
		//fmt.Println("@@@ temp", temp)
		//temp := int64(0)
		for right < n && nums[right] < queries[queriesIndex[j]] {
			temp += int64(queries[queriesIndex[j]] - nums[right] - (nums[right] - queries[queriesIndex[0]]))
			right++
		}
		preTemp = temp

		//fmt.Println(left, right, temp)

		ans[queriesIndex[j]] = ans[queriesIndex[0]] + int64((left-n+right+1)*(queries[queriesIndex[j]]-queries[queriesIndex[0]])) + temp
	}

	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
