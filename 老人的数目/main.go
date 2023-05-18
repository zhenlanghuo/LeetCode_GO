package main

import (
	"fmt"
	"strconv"
)

func main() {

}

func countSeniors(details []string) int {

	ans := 0

	for _, detail := range details {
		fmt.Println(detail[12:14])
		age, _ := strconv.ParseInt(detail[12:14], 10, 64)
		if age > 60 {
			ans++
		}
	}

	return ans
}
