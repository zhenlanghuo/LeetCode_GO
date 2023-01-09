package main

import "fmt"

func main() {
	fmt.Println(xorBeauty([]int{1, 4}))
	fmt.Println(xorBeauty([]int{15, 45, 20, 2, 34, 35, 5, 44, 32, 30}))
}

func xorBeauty(nums []int) int {
	result := 0
	//memory := make([][]int, len(nums))
	//for i := 0; i < len(nums); i++ {
	//	memory[i] = make([]int, len(nums))
	//	for j := 0; j < len(nums); j++ {
	//		memory[i][j] = -1
	//	}
	//}
	//var dfs func(level int, i, j, temp int)
	//dfs = func(level int, i, j, temp int) {
	//	if level == 2 {
	//		memory[i][j] = temp
	//		memory[j][i] = temp
	//		result = result ^ temp
	//		return
	//	}
	//
	//	for k := 0; k < len(nums); k++ {
	//		switch level {
	//		case 0:
	//			dfs(level+1, k, j, nums[k])
	//		case 1:
	//			dfs(level+1, i, k, temp&nums[k])
	//		}
	//	}
	//}

	//dfs(0, 0, 0, 0)

	for i := 0; i < len(nums); i++ {
		result = result ^ nums[i]
	}

	return result
}
