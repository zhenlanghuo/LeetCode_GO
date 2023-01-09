package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		result = append(result, stack[len(stack)-1].Val)
		temp := make([]*TreeNode, 0)
		for i := 0; i < len(stack); i++ {
			if stack[i].Left != nil {
				temp = append(temp, stack[i].Left)
			}
			if stack[i].Right != nil {
				temp = append(temp, stack[i].Right)
			}
		}
		stack = temp
	}

	return result
}
