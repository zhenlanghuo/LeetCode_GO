package main

import "container/heap"

func main() {

}

type State struct {
	id            int
	distFromStart int
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distFromStart < pq[j].distFromStart
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*State)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

type Pair struct {
	node   int
	weight int
}

type Graph struct {
	adj [][]*Pair
	//distTo [][]int
}

func Constructor(n int, edges [][]int) Graph {
	adj := make([][]*Pair, n)
	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		adj[from] = append(adj[from], &Pair{node: to, weight: weight})
	}
	//distTo := make([][]int, n)
	//for i := 0; i < n; i++ {
	//	distTo[i] = make([]int, n)
	//	for j := 0; j < n; j++ {
	//		if j == i {
	//			continue
	//		}
	//		distTo[i][j] = -1
	//	}
	//}
	return Graph{
		adj: adj,
		//distTo: distTo,
	}
}

func (this *Graph) AddEdge(edge []int) {
	from, to, weight := edge[0], edge[1], edge[2]
	this.adj[from] = append(this.adj[from], &Pair{node: to, weight: weight})
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {

	// 优先级队列，distFromStart 较小的排在前面
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &State{node1, 0})

	// 图中节点的个数
	V := len(this.adj)
	// 记录最短路径的权重，你可以理解为 dp table
	// 定义：distTo[i] 的值就是节点 start 到达节点 i 的最短路径权重
	distTo := make([]int, V)
	// 求最小值，所以 dp table 初始化为正无穷
	for i := range distTo {
		distTo[i] = -1
	}
	// base case，start 到 start 的最短距离就是 0
	distTo[node1] = 0

	for pq.Len() > 0 {
		curState := heap.Pop(&pq).(*State)
		curNodeID := curState.id
		curDistFromStart := curState.distFromStart

		if curNodeID == node2 {
			return curDistFromStart
		}

		if curDistFromStart > distTo[curNodeID] {
			// 已经有一条更短的路径到达 curNode 节点了
			continue
		}
		// 将 curNode 的相邻节点装入队列
		for _, nextNodePair := range this.adj[curNodeID] {
			nextNodeID, weight := nextNodePair.node, nextNodePair.weight
			// 看看从 curNode 达到 nextNode 的距离是否会更短
			distToNextNode := distTo[curNodeID] + weight
			if distTo[nextNodeID] < 0 || distTo[nextNodeID] > distToNextNode {
				// 更新 dp table
				distTo[nextNodeID] = distToNextNode
				// 将这个节点以及距离放入队列
				heap.Push(&pq, &State{nextNodeID, distToNextNode})
			}
		}
	}

	return -1
}
