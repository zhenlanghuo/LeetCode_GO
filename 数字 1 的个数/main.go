package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countDigitOne(122))
	fmt.Println(countDigitOneV1(122))
	fmt.Println(countDigitOne(3))
	fmt.Println(countDigitOneV1(3))
	fmt.Println(countDigitOne(1223))
	fmt.Println(countDigitOneV1(1223))
}

func countDigitOne(n int) (ans int) {
	// 数位DP

	// 将n转为字符串s
	s := strconv.FormatUint(uint64(n), 10)

	memory := make([][][2]int, len(s))
	for i := 0; i < len(s); i++ {
		memory[i] = make([][2]int, 1<<10)
		for j := 0; j < 1<<10; j++ {
			for k := 0; k < 2; k++ {
				memory[i][j][k] = -1
			}
		}
	}

	// dp函数表示构造从左到右第i位及其后面数位数字中数字1出现的个数
	// cnt表示目前数字1的个数
	// isLimit 前面一位数字是否已经达到对应数位的数字；为false的话, 当前数位可考虑的最大数字为9；为true的话，当前数位可考虑的最大数字为s[i]
	var dp func(i, cnt int, isLimit bool) int

	dp = func(i, cnt int, isLimit bool) int {
		if i == len(s) {
			return cnt
		}

		if memory[i][cnt][boolToInt(isLimit)] != -1 {
			return memory[i][cnt][boolToInt(isLimit)]
		}

		res := 0
		down := 0
		up := 9
		if isLimit {
			up = int(s[i] - '0')
		}
		for j := down; j <= up; j++ {
			if j == 1 {
				res += dp(i+1, cnt+1, isLimit && byte(j+'0') == s[i])
			} else {
				res += dp(i+1, cnt, isLimit && byte(j+'0') == s[i])
			}
		}

		memory[i][cnt][boolToInt(isLimit)] = res
		return res
	}

	return dp(0, 0, true)
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func countDigitOneV1(n int) (ans int) {
	// mulk 表示 10^k
	// 在下面的代码中，可以发现 k 并没有被直接使用到（都是使用 10^k）
	// 但为了让代码看起来更加直观，这里保留了 k
	for k, mulk := 0, 1; n >= mulk; k++ {
		ans += (n/(mulk*10))*mulk + min(max(n%(mulk*10)-mulk+1, 0), mulk)
		mulk *= 10
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
