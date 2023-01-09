package main

import "fmt"

func main() {
	fmt.Println(isPossible(4, [][]int{{1, 2}, {1, 3}, {1, 4}}))
}

func isPossible(n int, edges [][]int) bool {

	degrees := make([]int, n+1)
	edgeMap := make(map[int]map[int]bool)
	for _, edge := range edges {
		degrees[edge[0]]++
		degrees[edge[1]]++
		if _, ok := edgeMap[edge[0]]; !ok {
			edgeMap[edge[0]] = make(map[int]bool)
		}
		edgeMap[edge[0]][edge[1]] = true
		if _, ok := edgeMap[edge[1]]; !ok {
			edgeMap[edge[1]] = make(map[int]bool)
		}
		edgeMap[edge[1]][edge[0]] = true

	}

	count := 0
	nodes := make([]int, 0)
	for i := 1; i < n+1; i++ {
		if degrees[i]%2 != 0 {
			nodes = append(nodes, i)
			count++
		}
	}

	fmt.Println(edgeMap)
	fmt.Println(nodes)

	if count == 0 {
		return true
	}

	if count != 2 && count != 4 {
		return false
	}

	if count == 2 {
		if _, ok := edgeMap[nodes[0]][nodes[1]]; !ok {
			return true
		}

		for i := 1; i <= n; i++ {
			if i == nodes[0] || i == nodes[1] {
				continue
			}
			_, ok1 := edgeMap[i][nodes[0]]
			_, ok2 := edgeMap[i][nodes[1]]
			if !ok1 && !ok2 {
				return true
			}
		}
	}

	if count == 4 {
		_, ok1 := edgeMap[nodes[0]][nodes[1]]
		_, ok2 := edgeMap[nodes[2]][nodes[3]]
		if !ok1 && !ok2 {
			return true
		}

		_, ok1 = edgeMap[nodes[0]][nodes[2]]
		_, ok2 = edgeMap[nodes[1]][nodes[3]]
		if !ok1 && !ok2 {
			return true
		}

		_, ok1 = edgeMap[nodes[0]][nodes[3]]
		_, ok2 = edgeMap[nodes[1]][nodes[2]]
		if !ok1 && !ok2 {
			return true
		}
	}

	//fmt.Println(nodes)

	return false
}

//func permutations(nodes []int, level int) {
//	if level == len(nodes) {
//		fmt.Println(nodes)
//		temp := make([]int, len(nodes))
//		copy(temp, nodes)
//		result = append(result, temp)
//		return
//	}
//
//	for i := level; i < len(nodes); i++ {
//		nodes[level], nodes[i] = nodes[level], nodes[i]
//		permutations(nodes, level+1)
//		nodes[level], nodes[i] = nodes[level], nodes[i]
//	}
//}
