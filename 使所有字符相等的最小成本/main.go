package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumCost("0011"))
}

func minimumCost(s string) int64 {
	n := len(s)
	dp1 := make([][]int64, n+1)
	for i := 0; i < n+1; i++ {
		dp1[i] = make([]int64, 2)
	}

	bytes := []byte(s)
	for i := int64(1); i < int64(n)+1; i++ {
		switch bytes[i-1] {
		case '0':
			dp1[i][0] = min(dp1[i-1][0], dp1[i-1][1]+i-1)
			dp1[i][1] = min(dp1[i-1][0]+i, dp1[i-1][1]+i+i-1)
		case '1':
			dp1[i][0] = min(dp1[i-1][0]+i+i-1, dp1[i-1][1]+i)
			dp1[i][1] = min(dp1[i-1][0]+i-1, dp1[i-1][1])
		}
	}
	fmt.Println(dp1)

	dp2 := make([][]int64, n+1)
	for i := 0; i < n+1; i++ {
		dp2[i] = make([]int64, 2)
	}
	l, r := 0, len(bytes)-1
	for l < r {
		bytes[l], bytes[r] = bytes[r], bytes[l]
		l++
		r--
	}
	for i := int64(1); i < int64(n)+1; i++ {
		switch bytes[i-1] {
		case '0':
			dp2[i][0] = min(dp2[i-1][0], dp2[i-1][1]+i-1)
			dp2[i][1] = min(dp2[i-1][0]+i+1, dp2[i-1][1]+i+i-1)
		case '1':
			dp2[i][0] = min(dp2[i-1][0]+i+i-1, dp2[i-1][1]+i)
			dp2[i][1] = min(dp2[i-1][0]+i-1, dp2[i-1][1])
		}
	}

	l, r = 0, len(dp2)-1
	for l < r {
		dp2[l], dp2[r] = dp2[r], dp2[l]
		l++
		r--
	}
	fmt.Println(dp2)

	mn := int64(math.MaxInt64)
	for i := 1; i <= n; i++ {
		mn = min(mn, min(dp1[i][0]+dp2[i][0], dp1[i][1]+dp2[i][1]))
	}

	return mn
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
