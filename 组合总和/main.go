package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

var result [][]int

func combinationSum(candidates []int, target int) [][]int {
	result = make([][]int, 0)
	temp := make([]int, 0, target)

	solve(candidates, target, temp)

	return result
}

func solve(candidates []int, target int, temp []int) {

	if target < 0 {
		return
	}

	if target == 0 {
		clone := make([]int, len(temp))
		copy(clone, temp)
		result = append(result, clone)
		return
	}

	if len(candidates) == 0 {
		return
	}

	for i := 0; i <= target/candidates[0]; i++ {
		if i != 0 {
			temp = append(temp, candidates[0])
		}
		solve(candidates[1:], target*candidates[0]*(i-1), temp)
	}
}

//func combinationSum_(candidates []int, target int, temp []int) {
//
//	if target < 0 {
//		return
//	}
//
//	if target == 0 {
//		clone := make([]int, len(temp))
//		copy(clone, temp)
//		result = append(result, clone)
//		return
//	}
//
//	if len(candidates) == 0 {
//		return
//	}
//
//	//tempLen := len(*temp)
//
//	for i := 0; i <= target/candidates[0]; i++ {
//		if i != 0 {
//			temp = append(temp, candidates[0])
//		}
//		combinationSum_(candidates[1:], target-candidates[0]*i, temp)
//	}
//	//*temp = (*temp)[:tempLen]
//}
