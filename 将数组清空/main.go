package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(countOperationsToEmptyArray([]int{3, 4, -1}))
	fmt.Println(countOperationsToEmptyArray([]int{1, 2, 3}))
	fmt.Println(countOperationsToEmptyArray([]int{1, 2, 4, 3}))
	fmt.Println(countOperationsToEmptyArray([]int{4, 3, 1, 2, 5}))
	fmt.Println(countOperationsToEmptyArray([]int{6, 18, 13, -15}))
}

// 树状数组模板
type BIT []int

// 将下标 i 上的数加一
func (t BIT) inc(i int) {
	// -i <=> ~i+1
	for ; i < len(t); i += i & -i {
		t[i]++
	}
}

// 返回闭区间 [1, i] 的元素和
func (t BIT) sum(i int) int {
	sum := 0
	for ; i > 0; i &= i - 1 {
		sum += t[i]
	}
	return sum
}

// 返回闭区间 [left, right] 的元素和
func (t BIT) query(left, right int) int {
	return t.sum(right) - t.sum(left-1)
}

//func countOperationsToEmptyArray(nums []int) int64 {
//	n := len(nums)
//	id := make([]int, n)
//	for i := range id {
//		id[i] = i
//	}
//	sort.Slice(id, func(i, j int) bool { return nums[id[i]] < nums[id[j]] })
//
//	ans := n            // 先把 n 计入答案
//	t := make(BIT, n+1) // 下标从 1 开始
//	pre := 1            // 上一个最小值的位置，初始为 1
//	for k, i := range id {
//		i++           // 下标从 1 开始
//		if i >= pre { // 从 pre 移动到 i，跳过已经删除的数
//			ans += i - pre - t.query(pre, i)
//		} else { // 从 pre 移动到 n，再从 1 移动到 i，跳过已经删除的数
//			ans += n - pre + i - k + t.query(i, pre-1)
//		}
//		t.inc(i) // 删除 i
//		pre = i
//	}
//	return int64(ans)
//}

func countOperationsToEmptyArray(nums []int) int64 {
	n := len(nums)
	id := make([]int, n)
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool { return nums[id[i]] < nums[id[j]] })

	count := make([]int, 4*n)
	var build func(o, l, r int)
	build = func(o, l, r int) {
		if l == r {
			return
		}
		m := (l + r) / 2
		build(o*2, l, m)
		build(o*2+1, m+1, r)
		count[o] = count[o*2] + count[o*2+1]
	}

	var update func(o, l, r, idx int)
	update = func(o, l, r, idx int) {
		if l == r {
			count[o] = 1
			return
		}
		m := (l + r) / 2
		if idx <= m {
			update(o*2, l, m, idx)
		} else {
			update(o*2+1, m+2, r, idx)
		}
		count[o] = count[o*2] + count[o*2+1]
	}

	var query func(o, l, r, ql, qr int) int
	query = func(o, l, r, ql, qr int) int {
		if l >= ql && r <= qr {
			return count[o]
		}

		sum := 0
		m := (l + r) / 2

		if ql <= m {
			sum += query(o*2, l, m, ql, qr)
		}

		if qr > m {
			sum += query(o*2+1, m+1, r, ql, qr)
		}

		return sum
	}

	ans := n // 先把 n 计入答案
	//t := make(BIT, n+1) // 下标从 1 开始
	pre := 1 // 上一个最小值的位置，初始为 1
	for k, i := range id {
		i++           // 下标从 1 开始
		if i >= pre { // 从 pre 移动到 i，跳过已经删除的数
			ans += i - pre - query(0, 1, n, pre, i)
		} else { // 从 pre 移动到 n，再从 1 移动到 i，跳过已经删除的数
			ans += n - pre + i - k + query(0, 1, n, i, pre-1)
		}
		update(0, 1, n, i) // 删除 i
		pre = i
	}
	return int64(ans)
}
