package main

import "fmt"

func main() {
	fmt.Println(reverseWords("God Ding"))
}

func reverseWords(s string) string {
	bytes := []byte(s)

	start := 0
	for start < len(bytes) {
		if bytes[start] == byte(' ') {
			start++
			continue
		}

		end := start
		for end+1 < len(bytes) && bytes[end+1] != byte(' ') {
			end++
		}

		reverse(bytes, start, end)
		start = end + 1
	}

	return string(bytes)
}

func reverse(bytes []byte, start, end int) {
	for start < end {
		bytes[start], bytes[end] = bytes[end], bytes[start]
		start++
		end--
	}
}
