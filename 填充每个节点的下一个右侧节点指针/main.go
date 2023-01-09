package main

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil || root.Left == nil {
		return root
	}

	// 先把下一层的左右节点连起来
	root.Left.Next = root.Right

	if root.Next != nil {
		// 将右节点与上一层的隔壁节点的左节点连起来
		root.Right.Next = root.Next.Left
	}

	connect(root.Left)
	connect(root.Right)

	return root
}

//func connect(root *Node) *Node {
//
//	// 先遍历有多少层
//	tempRoot := root
//	level := 0
//	for tempRoot != nil {
//		level = level + 1
//		tempRoot = tempRoot.Left
//	}
//
//	record := make([][]*Node, level, level)
//	for i:=0;i<level;i++{
//		len_ := 1 << (uint(level)-1)
//		record[i] = make([]*Node, 0, len_)
//	}
//	connect_(root, 1, record)
//
//	for i:=0;i<level;i++{
//		for j:=0;j<len(record[i]);j++{
//			if j==len(record[i])-1 {
//				record[i][j].Next = nil
//				continue
//			}
//			record[i][j].Next = record[i][j+1]
//		}
//	}
//
//	return root
//}

//func connect_(root *Node, level int, record [][]*Node) {
//	if root == nil {
//		return
//	}
//
//	record[level] = append(record[level], root)
//	connect_(root.Left, level+1, record)
//	connect_(root.Right, level+1, record)
//}
