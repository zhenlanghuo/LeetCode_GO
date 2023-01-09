package main

import "strconv"

func main() {

}

func numDupDigitsAtMostN(n int) int {

	return n - countSpecialNumbers(n)

}

func countSpecialNumbers(n int) int {
	// 数位DP

	// 将n转为字符串s
	s := strconv.FormatUint(uint64(n), 10)

	memory := make([][][2][2]int, len(s))
	for i := 0; i < len(s); i++ {
		memory[i] = make([][2][2]int, 1<<10)
		for j := 0; j < 1<<10; j++ {
			for k := 0; k < 2; k++ {
				for h := 0; h < 2; h++ {
					memory[i][j][k][h] = -1
				}
			}
		}
	}

	// dp函数表示构造从左到右第i位及其后面数位数字的方案数
	// mask存储已经使用的数字（1~9）
	// isLimit 前面一位数字是否已经达到对应数位的数字；为false的话, 当前数位可考虑的最大数字为9；为true的话，当前数位可考虑的最大数字为s[i]
	// isNum 前面是否已经填过数字；为true的话，当前位必须填数字，0在选择范围内；为false的话，当前位可填可不填，填的话，0不在选择范围内
	var dp func(i, mask int, isLimit, isNum bool) int

	dp = func(i, mask int, isLimit, isNum bool) int {
		if i == len(s) {
			if isNum {
				return 1
			} else {
				return 0
			}
		}

		if memory[i][mask][boolToInt(isLimit)][boolToInt(isNum)] != -1 {
			return memory[i][mask][boolToInt(isLimit)][boolToInt(isNum)]
		}

		res := 0
		if !isNum {
			res += dp(i+1, mask, false, isNum)
		}
		down := 1
		if isNum {
			down = 0
		}
		up := 9
		if isLimit {
			up = int(s[i] - '0')
		}
		for j := down; j <= up; j++ {
			if mask&(1<<uint(j)) == 1<<uint(j) {
				continue
			}
			res += dp(i+1, mask^1<<uint(j), isLimit && byte(j+'0') == s[i], true)
		}

		memory[i][mask][boolToInt(isLimit)][boolToInt(isNum)] = res
		return res
	}

	return dp(0, 0, true, false)
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
