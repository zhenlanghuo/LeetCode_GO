package main

import "fmt"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result []*TreeNode
var record map[string]int

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	result = make([]*TreeNode, 0)
	record = make(map[string]int)
	dfs(root)
	return result
}

func dfs(root *TreeNode) string {

	if root == nil {
		return "none"
	}

	left := dfs(root.Left)
	right := dfs(root.Right)

	temp := fmt.Sprintf("%v,%v,%v", root.Val, left, right)

	if _, ok := record[temp]; !ok {
		record[temp] = 0
	}
	record[temp]++
	if record[temp] == 2 {
		result = append(result, root)
	}

	return temp
}
