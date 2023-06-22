package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	temp := head

	s1 := make([]int, 0)
	s2 := make([]int, 0)

	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}

	reverse(s1)
	reverse(s2)

	if len(s1) < len(s2) {
		s1, s2 = s2, s1
	}

	s3 := make([]int, 0)

	add := 0
	i := 0
	for ; i < len(s1) && i < len(s2); i++ {
		val := s1[i] + s2[i] + add
		add = val / 10
		s3 = append(s3, val%10)
	}

	for ; i < len(s1); i++ {
		val := s1[i] + add
		add = val / 10
		s3 = append(s3, val%10)
	}

	if add != 0 {
		s3 = append(s3, add)
	}

	reverse(s3)

	for i := 0; i < len(s3); i++ {
		temp.Next = &ListNode{Val: s3[i]}
		temp = temp.Next
	}

	return head.Next
}

func reverse(num []int) {
	l, r := 0, len(num)-1
	for l < r {
		num[l], num[r] = num[r], num[l]
		l++
		r--
	}
}
