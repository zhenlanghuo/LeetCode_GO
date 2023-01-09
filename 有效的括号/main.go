package main

import "fmt"

func main() {
	fmt.Println(isValid("[()]{[()]}{()}"))
}

var pairs = map[byte]byte{
	'(': ')',
	'{': '}',
	'[': ']',
}

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if _, ok := pairs[s[i]]; ok {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}

			if pairs[stack[len(stack)-1]] != s[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return true
	}

	return false
}
