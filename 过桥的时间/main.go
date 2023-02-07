package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

func main() {
	//fmt.Println(findCrossingTime(1, 3, [][]int{{1, 1, 2, 1}, {1, 1, 3, 1}, {1, 1, 4, 1}}))
	//fmt.Println(findCrossingTime(3, 2, [][]int{{1, 9, 1, 8}, {10, 10, 10, 10}}))
	fmt.Println(findCrossingTime(14, 7, [][]int{{8, 5, 9, 9}, {9, 10, 8, 8}, {9, 5, 4, 5}, {8, 8, 2, 2}, {8, 8, 9, 3}, {3, 6, 5, 6}, {4, 5, 6, 2}}))
}

type Pair struct {
	index int
	end   int
}

type waitHp []Pair

func (w waitHp) Len() int            { return len(w) }
func (w waitHp) Less(i, j int) bool  { return w[i].index > w[j].index }
func (w waitHp) Swap(i, j int)       { w[i], w[j] = w[j], w[i] }
func (w *waitHp) Push(x interface{}) { *w = append(*w, x.(Pair)) }
func (w *waitHp) Pop() interface{} {
	v := (*w)[len(*w)-1]
	*w = (*w)[:len(*w)-1]
	return v
}

type workHp []Pair

func (w workHp) Len() int            { return len(w) }
func (w workHp) Less(i, j int) bool  { return w[i].end < w[j].end }
func (w workHp) Swap(i, j int)       { w[i], w[j] = w[j], w[i] }
func (w *workHp) Push(x interface{}) { *w = append(*w, x.(Pair)) }
func (w *workHp) Pop() interface{} {
	v := (*w)[len(*w)-1]
	*w = (*w)[:len(*w)-1]
	return v
}

func findCrossingTime(n int, k int, time [][]int) int {
	// 下标越大, 效率越低
	sort.SliceStable(time, func(i, j int) bool {
		if time[i][0]+time[i][2] == time[j][0]+time[j][2] {
			return i < j
		}

		return time[i][0]+time[i][2] < time[j][0]+time[j][2]
	})

	waitL, waitR := make(waitHp, 0), make(waitHp, 0)
	workL, workR := make(workHp, 0), make(workHp, 0)
	for i := 0; i < len(time); i++ {
		waitL = append(waitL, Pair{index: i})
	}
	heap.Init(&waitL)
	heap.Init(&waitR)
	heap.Init(&workL)
	heap.Init(&workR)

	cur := 0
	newCount := 0
	oldCount := n
	for n > 0 {
		// fmt.Println(cur, newCount, oldCount)
		//fmt.Println(fmt.Sprintf("waitL: %v", waitL))
		//fmt.Println(fmt.Sprintf("waitR: %v", waitR))
		//fmt.Println(fmt.Sprintf("workL: %v", workL))
		//fmt.Println(fmt.Sprintf("workR: %v", workR))

		for len(workR) > 0 && workR[0].end <= cur {
			if oldCount > 0 {
				i := heap.Pop(&workR).(Pair).index
				heap.Push(&waitR, Pair{index: i})
				oldCount--
			}
		}

		for len(workL) > 0 && workL[0].end <= cur {
			i := heap.Pop(&workL).(Pair).index
			heap.Push(&waitL, Pair{index: i})
			newCount++
		}

		//fmt.Println(fmt.Sprintf("waitL: %v", waitL))
		//fmt.Println(fmt.Sprintf("waitR: %v", waitR))

		if len(waitR) > 0 {
			if oldCount == 0 && len(waitR) == 1 {
				i := waitR[0].index
				return cur + time[i][2]
			}

			// 过桥, 右到左
			i := heap.Pop(&waitR).(Pair).index
			heap.Push(&workL, Pair{index: i, end: cur + time[i][3] + time[i][2]})
			cur += time[i][2]
			continue
		} else if len(waitL) > 0 {
			if oldCount != 0 && oldCount != len(workR) {
				// 过桥, 左到右
				i := heap.Pop(&waitL).(Pair).index
				heap.Push(&workR, Pair{index: i, end: cur + time[i][1] + time[i][0]})
				cur += time[i][0]
				continue
			}
		}

		cur = math.MaxInt64
		if len(workR) != 0 {
			cur = workR[0].end
		}
		if len(workL) != 0 {
			cur = min(workL[0].end, cur)
		}
	}

	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
