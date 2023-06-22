package main

import "fmt"

func main() {
	fmt.Println(minLength("ABFCACDB"))
	fmt.Println(minLength("ACBBD"))
}

func minLength(s string) int {
	bytes := []byte(s)
	for {
		flag := true
		for i := 0; i < len(bytes)-1; i++ {
			if (bytes[i] == 'A' && bytes[i+1] == 'B') || (bytes[i] == 'C' && bytes[i+1] == 'D') {
				bytes = append(bytes[:i], bytes[i+2:]...)
				flag = false
			}
		}
		if flag {
			break
		}
	}

	return len(bytes)
}
