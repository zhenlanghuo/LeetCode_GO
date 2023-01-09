package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	stack := make([]*TreeNode, 0)
	stack = append(stack, &TreeNode{Val: nums[0]})
	for i := 1; i < len(nums); i++ {
		temp := &TreeNode{Val: nums[i]}
		for len(stack) != 0 && nums[i] > stack[len(stack)-1].Val {
			temp.Left = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		if len(stack) != 0 {
			stack[len(stack)-1].Right = temp
		}

		stack = append(stack, temp)
	}

	return stack[0]
}
