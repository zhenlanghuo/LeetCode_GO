package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(atMostNGivenDigitSet([]string{"1", "3", "5", "7"}, 100))
}

func atMostNGivenDigitSet(digits []string, n int) int {
	s := strconv.FormatUint(uint64(n), 10)

	memory := make([][2][2]int, len(s))
	for i := 0; i < len(s); i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				memory[i][j][k] = -1
			}
		}
	}

	var dp func(i int, isLimit, isNum bool) int

	// dp函数表示构造从左到右第i位及其后面数位数字的方案数
	// isLimit 表示前面一位数字是否已经达到对应数位的数字；为false的话, 当前数位可考虑的最大数字为9；为true的话，当前数位可考虑的最大数字为s[i]
	// isNum 前面是否已经填过数字；为true的话，当前位必须填数字，0在选择范围内；为false的话，当前位可填可不填，填的话，0不在选择范围内
	dp = func(i int, isLimit, isNum bool) int {
		if i == len(s) {
			if isNum {
				return 1
			}
			return 0
		}

		if memory[i][boolToInt(isLimit)][boolToInt(isNum)] != -1 {
			return memory[i][boolToInt(isLimit)][boolToInt(isNum)]
		}

		res := 0
		if !isNum {
			res += dp(i+1, false, isNum)
		}

		down := 1
		if isNum {
			down = 0
		}
		up := 9
		if isLimit {
			up = int(s[i] - '0')
		}
		for _, digit := range digits {
			digitInt := int(digit[0] - '0')
			if digitInt < down || digitInt > up {
				continue
			}
			res += dp(i+1, isLimit && digit[0] == s[i], true)
		}
		memory[i][boolToInt(isLimit)][boolToInt(isNum)] = res
		return res
	}

	return dp(0, true, false)
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
