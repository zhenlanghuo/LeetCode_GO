package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

type monotoneQueue struct {
	list *list.List
	nums []int
}

func (m *monotoneQueue) push(i int) {
	for m.list.Len() != 0 && m.nums[i] > m.nums[m.list.Back().Value.(int)] {
		m.list.Remove(m.list.Back())
	}
	m.list.PushBack(i)
}

func (m *monotoneQueue) top() int {
	return m.list.Front().Value.(int)
}

func (m *monotoneQueue) removeTop() {
	m.list.Remove(m.list.Front())
}

func maxSlidingWindow(nums []int, k int) []int {

	result := make([]int, 0, len(nums)-k+1)

	mq := &monotoneQueue{list: list.New(), nums: nums}
	for i := 0; i < k; i++ {
		mq.push(i)
	}
	result = append(result, nums[mq.top()])

	for i := k; i < len(nums); i++ {
		if mq.top() == i-k {
			mq.removeTop()
		}
		mq.push(i)
		result = append(result, nums[mq.top()])
	}

	return result
}
