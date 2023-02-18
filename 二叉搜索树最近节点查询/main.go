package main

import "fmt"

func main() {
	fmt.Println(findMax([]int{1, 2, 4, 6, 9, 13, 14, 15}, 2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestNodes(root *TreeNode, queries []int) [][]int {
	arr := make([]int, 0)
	var midTravel func(root *TreeNode)
	midTravel = func(root *TreeNode) {
		if root == nil {
			return
		}
		midTravel(root.Left)
		arr = append(arr, root.Val)
		midTravel(root.Right)
	}
	midTravel(root)

	fmt.Println(arr)
	ans := make([][]int, 0, len(queries))
	for _, q := range queries {
		ans = append(ans, []int{findMax(arr, q), findMin(arr, q)})
	}

	return ans
}

func findMax(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		fmt.Println(l, r, mid)
		if arr[mid] == target {
			return target
		} else if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if l-1 >= 0 {
		return arr[l-1]
	}
	return -1
}

func findMin(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == target {
			return target
		} else if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if l < len(arr) {
		return arr[l]
	}
	return -1
}
