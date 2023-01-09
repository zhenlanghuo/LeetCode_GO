package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || root == nil {
		return root
	}

	leftRoot := lowestCommonAncestor(root.Left, p, q)
	rightRoot := lowestCommonAncestor(root.Right, p, q)

	if leftRoot == nil && rightRoot == nil {
		return nil
	}

	if leftRoot != nil && rightRoot != nil {
		return root
	}

	if leftRoot != nil {
		return leftRoot
	}

	if rightRoot != nil {
		return rightRoot
	}

	return nil
}

//
//func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//	pList := find(root, p)
//	qList := find(root, q)
//
//	n := len(pList)
//	if n > len(qList) {
//		n = len(qList)
//	}
//
//	var result *TreeNode
//	for i := 0; i < n; i++ {
//		if pList[i] != qList[i] {
//			break
//		}
//		result = pList[i]
//	}
//
//	return result
//}

func find(root, target *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}

	if root == target {
		return []*TreeNode{root}
	}

	result := find(root.Left, target)
	if result != nil {
		result = append([]*TreeNode{root}, result...)
		return result
	}

	result = find(root.Right, target)
	if result != nil {
		result = append([]*TreeNode{root}, result...)
		return result
	}

	return nil
}
