package main

import (
	"fmt"
	"sort"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	sums := make([]int64, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		temp := make([]*TreeNode, 0)
		sum := int64(0)
		for _, node := range queue {
			sum += int64(node.Val)
			if node.Left != nil {
				temp = append(temp, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
			}
		}
		sums = append(sums, sum)
		queue = temp
	}

	fmt.Println(sums)

	if len(sums) < k {
		return -1
	}

	sort.Slice(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})

	return sums[k-1]
}
