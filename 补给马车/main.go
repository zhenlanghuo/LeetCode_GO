package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(supplyWagon([]int{7, 3, 6, 1, 8}))
	fmt.Println(supplyWagon([]int{1, 3, 1, 5}))
}

func supplyWagon(supplies []int) []int {
	n := len(supplies)
	targetN := n / 2
	l := list.New()
	for _, supply := range supplies {
		l.PushBack(supply)
	}

	for l.Len() != targetN {
		head := l.Front()
		mn := head.Value.(int) + head.Next().Value.(int)
		mnStartNode := head
		for head.Next().Next() != nil {
			head = head.Next()
			temp := head.Value.(int) + head.Next().Value.(int)
			if temp < mn {
				mn = temp
				mnStartNode = head
			}
		}
		mnStartNode.Value = mn
		l.Remove(mnStartNode.Next())
	}

	ans := make([]int, targetN)
	for i := 0; i < targetN; i++ {
		ans[i] = l.Front().Value.(int)
		l.Remove(l.Front())
	}

	return ans
}
