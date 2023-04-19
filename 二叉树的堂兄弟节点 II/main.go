package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {

	faMap := make(map[*TreeNode]*TreeNode)
	handled := make(map[*TreeNode]bool)
	q := []*TreeNode{root}
	faMap[root] = nil
	for len(q) != 0 {
		sum := 0
		temp := q
		q = make([]*TreeNode, 0)
		for _, node := range temp {
			sum += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
				faMap[node.Left] = node
			}
			if node.Right != nil {
				q = append(q, node.Right)
				faMap[node.Right] = node
			}
		}

		for _, node := range temp {
			if handled[node] {
				continue
			}

			tempSum := 0
			if faMap[node] == nil {
				node.Val = 0
			} else {
				if faMap[node].Left != nil {
					tempSum += faMap[node].Left.Val
				}
				if faMap[node].Right != nil {
					tempSum += faMap[node].Right.Val
				}
				if faMap[node].Left != nil {
					faMap[node].Left.Val = sum - tempSum
					handled[faMap[node].Left] = true
				}
				if faMap[node].Right != nil {
					faMap[node].Right.Val = sum - tempSum
					handled[faMap[node].Right] = true
				}
			}
		}
	}

	return root
}
