package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(totalCost([]int{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3, 4))
	fmt.Println(totalCost([]int{1, 2, 4, 1}, 3, 3))
}

type Worker struct {
	cost  int
	index int
}

type Workers []*Worker

func (w Workers) Len() int { return len(w) }
func (w Workers) Less(i, j int) bool {
	if w[i].cost == w[j].cost {
		return w[i].index < w[j].index
	}
	return w[i].cost < w[j].cost
}
func (w Workers) Swap(i, j int)       { w[i], w[j] = w[j], w[i] }
func (w *Workers) Push(x interface{}) { *w = append(*w, x.(*Worker)) }
func (w *Workers) Pop() interface{} {
	v := (*w)[len(*w)-1]
	*w = (*w)[:len(*w)-1]
	return v
}

func totalCost(costs []int, k int, candidates int) int64 {
	n := len(costs)
	l, r := 0, n-1
	minHeap := make(Workers, 0)
	for l <= r && l < candidates {
		minHeap = append(minHeap, &Worker{index: l, cost: costs[l]})
		l++
		if l <= r {
			minHeap = append(minHeap, &Worker{index: r, cost: costs[r]})
			r--
		}
	}
	heap.Init(&minHeap)

	ans := int64(0)
	for i := 0; i < k; i++ {
		top := heap.Pop(&minHeap).(*Worker)
		ans += int64(top.cost)

		if l <= r {
			if top.index <= l {
				heap.Push(&minHeap, &Worker{index: l, cost: costs[l]})
				l++
			}
			if top.index >= r {
				heap.Push(&minHeap, &Worker{index: r, cost: costs[r]})
				r--
			}
		}
	}

	return ans
}
