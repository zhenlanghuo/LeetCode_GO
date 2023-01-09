package main

import "fmt"

func main() {
	fmt.Println(findIntegers(974367147))
}

func findIntegers(n int) int {

	bs := nToBinaryString(n)
	//fmt.Println(bs)

	memory := make([][2][2]int, len(bs))
	for i := 0; i < len(bs); i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				memory[i][j][k] = -1
			}
		}
	}

	// isOneBefore表示前1位是1
	// isLimit 表示前面一位数字是否已经达到对应数位的数字：为false的话, 当前数位可考虑的最大数字为1；为true的话，当前数位可考虑的最大数字为s[i]
	var dp func(i int, isOneBefore, isLimit bool) int

	dp = func(i int, isOneBefore, isLimit bool) int {
		if i == len(bs) {
			return 1
		}

		if memory[i][boolToInt(isOneBefore)][boolToInt(isLimit)] != -1 {
			return memory[i][boolToInt(isOneBefore)][boolToInt(isLimit)]
		}

		res := 0

		down := 0
		up := 1
		if isLimit {
			up = int(bs[i] - '0')
		}

		for j := down; j <= up; j++ {
			switch j {
			case 0:
				res += dp(i+1, false, isLimit && byte(j+'0') == bs[i])
			case 1:
				if !isOneBefore {
					res += dp(i+1, true, isLimit && byte(j+'0') == bs[i])
				}
			}
		}

		memory[i][boolToInt(isOneBefore)][boolToInt(isLimit)] = res
		return res
	}

	return dp(0, false, true)
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func nToBinaryString(n int) string {
	bytes := make([]byte, 0)

	for n > 0 {
		bytes = append(bytes, byte(n%2)+'0')
		n = n >> 1
	}
	reverse(bytes)

	return string(bytes)
}

func reverse(bytes []byte) {
	start, end := 0, len(bytes)-1
	for start < end {
		bytes[start], bytes[end] = bytes[end], bytes[start]
		start++
		end--
	}
}
