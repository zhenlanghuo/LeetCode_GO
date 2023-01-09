package main

import "fmt"

func main() {
	fmt.Println(intervalIntersection([][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}))
}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	result := make([][]int, 0)
	fi, si := 0, 0
	for fi < len(firstList) && si < len(secondList) {
		temp := intersection(firstList[fi], secondList[si])
		if temp != nil {
			result = append(result, temp)
		}
		if firstList[fi][1] > secondList[si][1] {
			si++
		} else {
			fi++
		}
	}

	return result
}

func intersection(x []int, y []int) []int {
	result := []int{max(x[0], y[0]), min(x[1], y[1])}

	if result[0] > result[1] {
		return nil
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
