package 二叉树的锯齿形层序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	stack := make([]*TreeNode, 0)
	count := 0
	stack = append(stack, root)
	result := make([][]int, 0)
	for len(stack) != 0 {
		temp := make([]*TreeNode, 0)
		tempResult := make([]int, 0)
		for i := len(stack) - 1; i >= 0; i-- {
			temp = append(temp, stack[i])
			tempResult = append(tempResult, stack[i].Val)
		}
		stack = stack[:0]
		result = append(result, tempResult)
		for _, node := range temp {
			switch count % 2 {
			case 0:
				pushNode(&stack, node.Left)
				pushNode(&stack, node.Right)
			case 1:
				pushNode(&stack, node.Right)
				pushNode(&stack, node.Left)
			}
		}
		count++
	}

	return result
}

func pushNode(stack *[]*TreeNode, node *TreeNode) {
	if node == nil {
		return
	}
	*stack = append(*stack, node)
}
