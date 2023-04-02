package main

import "fmt"

func main() {
	//fmt.Println(findLengthOfShortestSubarray([]int{1, 2, 3, 10, 4, 2, 3, 5}))
	fmt.Println(findLengthOfShortestSubarray([]int{16, 10, 0, 3, 22, 1, 14, 7, 1, 12, 15}))
}

func findLengthOfShortestSubarray(arr []int) int {

	n := len(arr)
	leftMax := 0
	for leftMax+1 < n && arr[leftMax+1] >= arr[leftMax] {
		leftMax++
	}

	rightMin := n - 1
	for rightMin-1 >= 0 && arr[rightMin-1] <= arr[rightMin] {
		rightMin--
	}

	//fmt.Println(leftMax, rightMin)

	if rightMin <= leftMax {
		return 0
	}

	fmt.Println(leftMax, rightMin)
	fmt.Println(arr[:leftMax+1], arr[rightMin:])

	ans := max(leftMax+1, n-rightMin)
	left, right := 0, rightMin
	for left <= leftMax {
		for right < n && arr[right] < arr[left] {
			right++
		}
		fmt.Println(left, right, left+1+n-right)

		ans = max(ans, left+1+n-right)
		left++
	}

	return n - ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
