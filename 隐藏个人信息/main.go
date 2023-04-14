package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(maskPII("LeetCode@LeetCode.com"))
	fmt.Println(maskPII("AB@qq.com"))
	fmt.Println(maskPII("11(234)567-890"))
}

func maskPII(s string) string {

	if strings.Contains(s, "@") {
		return handleEmail(s)
	}

	return handlePhoneNumber(s)
}

func handleEmail(s string) string {

	splits := strings.Split(s, "@")

	name := strings.ToLower(splits[0])

	return fmt.Sprintf("%s*****%s@%s", name[0:1], name[len(name)-1:], strings.ToLower(splits[1]))
}

func handlePhoneNumber(s string) string {

	numCount := 0
	last4Num := make([]byte, 4)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i]-'0' <= 9 {
			if numCount <= 3 {
				last4Num[numCount] = s[i]
			}
			numCount++
		}
	}
	reverse(last4Num)

	nationNum := make([]byte, numCount-10)
	for i := 0; i < numCount-10; i++ {
		nationNum[i] = '*'
	}

	if numCount-10 > 0 {
		return fmt.Sprintf("+%s-***-***-%s", string(nationNum), string(last4Num))
	}

	return fmt.Sprintf("***-***-%s", string(last4Num))
}

func reverse(arr []byte) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}
