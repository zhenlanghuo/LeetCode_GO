package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

var (
	result []string
)

func generateParenthesis(n int) []string {
	result = make([]string, 0)
	solve(n, n, "")
	return result
}

func solve(left, right int, str string) {
	if left == 0 {
		for i := 0; i < right; i++ {
			str += ")"
		}
		result = append(result, str)
	} else if left == right {
		solve(left-1, right, str+"(")
	} else {
		solve(left-1, right, str+"(")
		solve(left, right-1, str+")")
	}
}
