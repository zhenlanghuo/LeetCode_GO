package main

import "fmt"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeQueries(root *TreeNode, queries []int) []int {
	m := make(map[int][]int)
	fatherMap := make(map[int]*TreeNode)
	nodeMap := make(map[int]*TreeNode)

	var travel func(root, fa *TreeNode) int
	travel = func(root, fa *TreeNode) int {
		if root == nil {
			return 0
		}

		nodeMap[root.Val] = root
		fatherMap[root.Val] = fa
		lh := travel(root.Left, root)
		rh := travel(root.Right, root)
		m[root.Val] = []int{lh, rh, max(lh, rh) + 1}
		return max(lh, rh) + 1
	}
	travel(root, nil)

	fmt.Println(m)
	fmt.Println(fatherMap)

	var check func(root *TreeNode) int
	check = func(root *TreeNode) int {
		fa := fatherMap[root.Val]
		if fa == nil {
			return max(m[root.Val][0], m[root.Val][1]) + 1
		}
		if fa.Left != nil && fa.Left.Val == root.Val {
			// 左子树
			oh := m[fa.Val][0]
			m[fa.Val][0] = max(m[root.Val][0], m[root.Val][1]) + 1
			ans := check(fa)
			m[fa.Val][0] = oh
			return ans
		} else {
			// 右子树
			oh := m[fa.Val][1]
			m[fa.Val][1] = max(m[root.Val][0], m[root.Val][1]) + 1
			ans := check(fa)
			m[fa.Val][1] = oh
			return ans
		}
	}

	ans := make([]int, 0, len(queries))
	for _, q := range queries {
		ol, or := m[q][0], m[q][1]
		m[q][0], m[q][1] = 0, 0
		ans = append(ans, check(nodeMap[q]))
		m[q][0], m[q][1] = ol, or
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
