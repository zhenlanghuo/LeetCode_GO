package main

import "fmt"

func main() {
	fmt.Println(validPalindrome("aabbac"))
}

func validPalindrome(s string) bool {
	bytes := []byte(s)
	return validPalindrome_(bytes, 0, len(bytes)-1, false)
}

func validPalindrome_(bytes []byte, start, end int, flag bool) bool {
	for start < end {
		if bytes[start] != bytes[end] {
			if flag {
				return false
			}
			return validPalindrome_(bytes, start+1, end, true) || validPalindrome_(bytes, start, end-1, true)
		}

		start++
		end--
	}

	return true
}
