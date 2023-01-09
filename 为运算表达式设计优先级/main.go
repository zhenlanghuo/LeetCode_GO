package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(diffWaysToCompute("2*3-4*5"))
}

var opMap = map[byte]bool{'+': true, '-': true, '*': true}

var dp [][][]int

func diffWaysToCompute(expression string) []int {
	splits := make([]string, 0)
	start := 0
	for start < len(expression) {
		temp := start
		for temp < len(expression) && !opMap[expression[temp]] {
			temp++
		}
		splits = append(splits, expression[start:temp])
		if temp < len(expression) {
			splits = append(splits, expression[temp:temp+1])
		}
		start = temp + 1
	}

	dp = make([][][]int, len(splits))
	for i := 0; i < len(splits); i++ {
		dp[i] = make([][]int, len(splits))
	}

	for i := 0; i < len(splits); i = i + 2 {
		num, _ := strconv.ParseInt(splits[i], 10, 32)
		dp[i][i] = []int{int(num)}
	}

	for ln := 3; ln <= len(splits); ln = ln + 2 {
		for l, r := 0, ln-1;  r < len(splits); {
			for k := l + 1; k < r; k = k + 2 {
				dp[l][r] = append(dp[l][r], calculate(dp[l][k-1], dp[k+1][r], splits[k])...)
				//fmt.Println(l, r, dp[l][r])
			}
			l = l + 2
			r = r + 2
		}
	}

	//fmt.Println(dp)

	//fmt.Println(splits)

	//diffWaysToCompute_(splits, 0, len(splits)-1)
	//fmt.Println(dp)

	return dp[0][len(splits)-1]
}

func diffWaysToCompute_(splits []string, start, end int) []int {
	if start > end {
		return []int{}
	}

	if dp[start][end] != nil {
		return dp[start][end]
	}

	if start == end {
		num, _ := strconv.ParseInt(splits[start], 10, 32)
		dp[start][end] = []int{int(num)}
		return dp[start][end]
	}

	//resultMap := make(map[int]bool)
	for i := start + 1; i <= end; i = i + 2 {
		leftNums := diffWaysToCompute_(splits, start, i-1)
		rightNums := diffWaysToCompute_(splits, i+1, end)
		temp := calculate(leftNums, rightNums, splits[i])
		dp[start][end] = append(dp[start][end], temp...)
	}

	return dp[start][end]
}

func calculate(aList, bList []int, op string) []int {
	result := make([]int, 0, len(aList)*len(bList))
	for _, a := range aList {
		for _, b := range bList {
			switch op {
			case "-":
				result = append(result, a-b)
			case "+":
				result = append(result, a+b)
			case "*":
				result = append(result, a*b)
			}
		}
	}
	return result
}

func mapToSlice(m map[int]bool) []int {
	result := make([]int, 0, len(m))
	for num := range m {
		result = append(result, num)
	}
	return result
}
