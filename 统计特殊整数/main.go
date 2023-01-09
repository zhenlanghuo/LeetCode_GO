package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(countSpecialNumbers(135))
	//fmt.Println(countSpecialNumbers(20))
	fmt.Println(countSpecialNumbers(9863))
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

//func countSpecialNumbers(n int) int {
//	if n < 10 {
//		return n
//	}
//	bytes := make([]byte, 0)
//	for n > 0 {
//		bytes = append(bytes, byte(n%10))
//		n = n / 10
//	}
//	reverse(bytes)
//	fmt.Println(bytes)
//
//	// n是一个m位数, 先求1~m-1位数[1~(1em)-1]的特殊整数个数
//	// 1位数的特殊整数个数为9
//	// 2位数的特殊整数个数为9*9（因为第一位不能有0，所以第一位选择只有9个数）
//	// 3位数的特殊整数个数为9*9*8
//	// 4位数的特殊整数个数为9*9*8*7
//	// 如此类推
//	sum := 0
//	cur := 0
//	for i := 0; i < len(bytes)-1; i++ {
//		if i == 0 {
//			cur = 9
//			sum += cur
//		} else {
//			cur = cur * (9 - i + 1)
//			sum += cur
//		}
//	}
//
//	fmt.Println(sum)
//
//	// 求[1em, n]的特殊整数个数
//	// 假设用 abcd 表示 n
//	// 先求第一位为 [1, a-1] 时的特殊整数个数
//	// 当第一位为a时, 求第二位为[0, b-1]时的特殊整数个数
//	// 当第二位为b时, 求第三位为[0, c-1]时的特殊整数个数
//	// 如此类推
//	used := make([]byte, 10)
//	for i := 0; i < len(bytes); i++ {
//		if i < len(bytes)-1 {
//			// 不是最高位时, num可选的范围是[0, bytes[i]-1], 可选的个数为 bytes[i]
//			num := int(bytes[i])
//			fmt.Println(i, num, used)
//			// 排除前面位数已经确定的数
//			for j := 0; j < int(bytes[i]); j++ {
//				if used[j] != 0 {
//					num--
//				}
//			}
//			// 如果是最高位, num可选的范围是[1, bytes[i]-1], 可选的个数为 bytes[i]-1
//			if i == 0 {
//				num = int(bytes[i]) - 1
//			}
//			fmt.Println(i, num)
//			used[bytes[i]]++
//			for j := 0; j < len(bytes)-1-i; j++ {
//				num = num * (9 - i - j)
//			}
//			sum += num
//			// 如果已经出现了两个数位为相同的数字, 就不用再继续了
//			if used[bytes[i]] == 2 {
//				break
//			}
//		} else {
//			fmt.Println(used)
//			for k := 0; k <= int(bytes[i]); k++ {
//				if used[k] == 0 {
//					sum++
//				}
//			}
//		}
//		fmt.Println(i, sum)
//	}
//
//	return sum
//}
//
//func reverse(bytes []byte) {
//	start, end := 0, len(bytes)-1
//	for start < end {
//		bytes[start], bytes[end] = bytes[end], bytes[start]
//		start++
//		end--
//	}
//}
