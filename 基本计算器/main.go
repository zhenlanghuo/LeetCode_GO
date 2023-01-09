package main

import "fmt"

func main() {
	//fmt.Println(calculate("(1-2) *4 /2"))
	//fmt.Println(calculate("-(1+(4+5+2)-3)+(6+8)"))
	fmt.Println(calculate("22"))
}

func calculate(s string) int {

	var re func(s string, start int) (int, int)
	re = func(s string, start int) (int, int) {
		num := 0
		sign := '+'
		stack := make([]int, 0)
		i := start
		for ; i < len(s); i++ {
			b := s[i]

			if b == '(' {
				num, i = re(s, i+1)
				//fmt.Println(num)
			}

			if isByteDigit(byte(b)) {
				num = num*10 + int(b-'0')
			}

			if isByteOps(byte(b)) || i == len(s)-1 || b==')' {
				//fmt.Println(start, i, stack, num, string(sign))
				switch sign {
				case '-':
					stack = append(stack, -num)
				case '+':
					stack = append(stack, num)
				case '*':
					stack[len(stack)-1] *= num
				case '/':
					stack[len(stack)-1] /= num
				}
				//fmt.Println(start, i, stack, num, string(sign))
				sign = int32(b)
				num = 0
			}

			if b == ')' {
				break
			}
		}

		//fmt.Println(start, stack)
		result := 0
		for _, v := range stack {
			result += v
		}

		return result, i
	}

	result, _ := re(s, 0)
	return result
}

func isByteDigit(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}

	return false
}

func isByteOps(b byte) bool {
	switch b {
	case '+', '-', '*', '/':
		return true
	}
	return false
}
