package main

import "fmt"

func main() {
	s := "((()))(()))"
	fmt.Print(longestValidParentheses(s))
}

func longestValidParentheses(s string) int {
	dp := make([]int, len(s))
	max := 0

	leftByte := []byte("(")[0]
	rightByte := []byte(")")[0]

	for i := 0; i < len(s); i++ {
		if s[i] != rightByte {
			continue
		}

		if i-1 >= 0 && s[i-1] == leftByte {
			if i-2 >= 0 {
				dp[i] = dp[i-2] + 2
			} else {
				dp[i] = 2
			}
		}

		if i-1 >= 0 && dp[i-1] > 0 && i-1-dp[i-1] >= 0 && s[i-1-dp[i-1]] == leftByte {
			dp[i] = dp[i-1] + 2
		}

		if i-dp[i] >= 0 && dp[i-dp[i]] > 0 {
			dp[i] = dp[i-dp[i]] + dp[i]
		}

		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}

//func longestValidParentheses(s string) int {
//	stack := make([]int, 0)
//	stack = append(stack, -1)
//	maxSum := 0
//	for i := 0; i < len(s); i++ {
//		if s[i] == '(' {
//			stack = append(stack, i)
//		} else {
//			stack = stack[:len(stack)-1]
//			if len(stack) != 0 {
//				maxSum = max(maxSum, i-stack[len(stack)-1])
//			} else {
//				stack = append(stack, i)
//
//			}
//		}
//	}
//
//	return maxSum
//}
//
//func max(i int, j int) int {
//	if i < j {
//		return j
//	}
//	return i
//}
