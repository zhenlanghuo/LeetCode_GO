package main

func main() {
	print()
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)

	dfs(root, 0, &result)

	return result
}

func dfs(root *TreeNode, level int, result *[][]int) {
	if root == nil {
		return
	}

	if len(*result) < level+1 {
		*result = append(*result, make([]int, 0))
	}
	(*result)[level] = append((*result)[level], root.Val)

	dfs(root.Left, level+1, result)
	dfs(root.Right, level+1, result)
}
