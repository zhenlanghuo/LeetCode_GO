package main

func main() {

}

func handleQuery(nums1 []int, nums2 []int, queries [][]int) []int64 {
	n := len(nums1)
	count := make([]int, 4*n)
	flip := make([]bool, 4*n)

	maintain := func(o int) {
		count[o] = count[o*2] + count[o*2+1]
	}

	var build func(o, l, r int)
	build = func(o, l, r int) {
		if l == r {
			count[o] = nums1[l-1]
			return
		}

		m := (l + r) / 2
		build(o*2, l, m)
		build(o*2+1, m+1, r)
		maintain(o)
	}

	do := func(o, l, r int) {
		count[o] = r - l + 1 - count[o]
		flip[o] = !flip[o]
	}

	var update func(o, l, r, L, R int)
	update = func(o, l, r, L, R int) {
		if l >= L && r <= R {
			do(o, l, r)
			return
		}

		m := (l + r) / 2

		if flip[o] {
			do(o*2, l, m)
			do(o*2+1, m+1, r)
			flip[o] = false
		}

		if m >= L {
			update(o*2, l, m, L, R)
		}
		if m+1 <= R {
			update(o*2+1, m+1, r, L, R)
		}
		maintain(o)
	}

	sum2 := int64(0)
	for _, num := range nums2 {
		sum2 += int64(num)
	}

	build(1, 1, n)
	ans := make([]int64, 0, len(queries))
	for _, q := range queries {
		if q[0] == 1 {
			update(1, 1, n, q[1]+1, q[2]+1)
		}

		if q[0] == 2 {
			sum2 += int64(count[1] * q[1])
		}

		if q[0] == 3 {
			ans = append(ans, sum2)
		}
	}

	return ans
}
