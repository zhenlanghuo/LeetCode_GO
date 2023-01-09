package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValid(root, nil, nil)
}

func isValid(root *TreeNode, max, min *int) bool {
	if root == nil {
		return true
	}

	if max != nil && root.Val >= *max {
		return false
	}

	if min != nil && root.Val <= *min {
		return false
	}

	return isValid(root.Left, &root.Val, min) && isValid(root.Right, max, &root.Val)
}

















//func isValid(root *TreeNode, max, min *int) bool {
//	if root == nil {
//		return true
//	}
//
//	if max != nil && root.Val >= *max {
//		return false
//	}
//
//	if min != nil && root.Val <= *min {
//		return false
//	}
//
//	return isValid(root.Left, &root.Val, nil) && isValid(root.Right, nil, &root.Val)
//}
//
//func solve(root *TreeNode) (bool, int, int) {
//	var (
//		leftMin, leftMax, rightMin, rightMax *int
//	)
//
//	if root.Left != nil {
//		isLeftValidBST, min, max := solve(root.Left)
//		if !isLeftValidBST {
//			return false, 0, 0
//		}
//
//		if root.Val <= root.Left.Val {
//			return false, 0, 0
//		}
//
//		if root.Val <= max {
//			return false, 0, 0
//		}
//
//		leftMin = &min
//		leftMax = &max
//	}
//
//	if root.Right != nil {
//		isRightValidBST, min, max := solve(root.Right)
//		if !isRightValidBST {
//			return false, 0, 0
//		}
//
//		if root.Val >= root.Right.Val {
//			return false, 0, 0
//		}
//
//		if root.Val >= min {
//			return false, 0, 0
//		}
//
//		rightMin = &min
//		rightMax = &max
//	}
//
//	min := root.Val
//	if rightMin != nil && min > *rightMin {
//		min = *rightMin
//	}
//	if leftMin != nil && min > *leftMin {
//		min = *leftMin
//	}
//
//	max := root.Val
//	if rightMax != nil && max < *rightMax {
//		max = *rightMax
//	}
//	if leftMax != nil && max < *leftMax {
//		max = *leftMax
//	}
//
//	return true, min, max
//}
