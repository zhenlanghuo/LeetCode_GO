package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(decodeString("3[a2[c]]"))
}

var (
	bytea            = []byte("a")[0]
	bytez            = []byte("z")[0]
	byteLeftBracket  = []byte("[")[0]
	byteRightBracket = []byte("]")[0]
)

func decodeString(s string) string {
	return string(decodeString_([]byte(s)))
}

func decodeString_(s []byte) []byte {
	result := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] >= bytea && s[i] <= bytez {
			result = append(result, s[i])
			continue
		}

		j := i
		for j < len(s) && s[j] != byteLeftBracket {
			j++
		}

		n, _ := strconv.ParseUint(string(s[i:j]), 10, 64)

		leftCount := 0
		rightCount := 0
		k := j
		for {
			if s[k] == byteLeftBracket {
				leftCount++
			}

			if s[k] == byteRightBracket {
				rightCount++
			}

			if rightCount == leftCount {
				break
			}
			k++
		}

		for n > 0 {
			result = append(result, decodeString_(s[j+1:k])...)
			n--
		}

		i = k
		fmt.Println(j, k, leftCount, rightCount)
	}

	return result
}
