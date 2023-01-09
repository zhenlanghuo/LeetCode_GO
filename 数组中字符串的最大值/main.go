package main

import "fmt"

func main() {
	fmt.Println(maximumValue([]string{"alic3","bob","3","4","00100"}))
	fmt.Println(maximumValue([]string{"1","01","001","0001"}))
}

func maximumValue(strs []string) int {
	result := 0
	for _, str := range strs {
		onlyNum := true
		for _, b := range str {
			if b-'0' >= 0 && b-'0' <= 9 {
				continue
			}
			onlyNum = false
		}
		if onlyNum {
			temp := 0
			for _, b := range str {
				temp = temp*10 + int(b-'0')
			}
			result = max(result, temp)
		} else {
			result = max(result, len(str))
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
