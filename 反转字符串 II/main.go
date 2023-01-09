package main

import "fmt"

func main() {
	fmt.Println(reverseStr("a", 2))
}

func reverseStr(s string, k int) string {
	bytes := []byte(s)

	start := 0
	for start < len(s) {
		if start+2*k-1 < len(s) {
			reverse(bytes, start, start+k-1)
			start += 2 * k
			continue
		}

		if (len(s) - start + 1) < k {
			reverse(bytes, start, len(s)-1)
		} else {
			end := start + k - 1
			if end >= len(s) {
				end = len(s) - 1
			}
			reverse(bytes, start, end)
		}
		break
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
