package main

import "fmt"

func main() {
	fmt.Println(baseNeg2(3))
	fmt.Println(baseNeg2(4))
}

func baseNeg2(n int) string {

	if n == 0 {
		return "0"
	}

	ans := make([]byte, 0)
	for n != 0 {
		if n%(-2) == 0 {
			ans = append(ans, '0')
		} else {
			ans = append(ans, '1')
			n -= 1
		}
		n = n / -2
	}

	l, r := 0, len(ans)-1
	for l < r {
		ans[l], ans[r] = ans[r], ans[l]
		l++
		r--
	}

	return string(ans)
}
