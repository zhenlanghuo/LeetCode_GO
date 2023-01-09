package main

import "fmt"

func main() {
	fmt.Println(minInsertions("(()))"))
	fmt.Println(minInsertions("())"))
	fmt.Println(minInsertions("))())("))
	fmt.Println(minInsertions("(((((("))
}

func minInsertions(s string) int {
	leftCount := 0
	result := 0
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b == ')' {
			if i == len(s)-1 || s[i+1] != ')' {
				result++
			} else {
				i++
			}
			if leftCount == 0 {
				result++
			} else {
				leftCount--
			}
		} else {
			leftCount++
		}
	}

	return result + leftCount*2
}
