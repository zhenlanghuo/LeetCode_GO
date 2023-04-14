package main

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {

	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	n := len(nums)
	ans := make([]int, n)

	type pair struct{ v, i int }
	stack := make([]pair, 0, n)

	for i := 0; i < n; i++ {

		for len(stack) != 0 && stack[len(stack)-1].v < nums[i] {
			ans[stack[len(stack)-1].i] = nums[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, pair{v: nums[i], i: i})
	}

	return ans
}
