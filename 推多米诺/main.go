package main

import "fmt"

func main() {
	fmt.Println(pushDominoes(".L.R...LR..L.."))
}

func pushDominoes(dominoes string) string {

	ans := make([]byte, len(dominoes))
	lastR := -1
	lastRL := -1

	for i, d := range dominoes {
		ans[i] = byte(d)

		if d == '.' {
			continue
		}

		if d == 'L' {
			if lastR != -1 {
				lastRL = i
				l, r := lastR, i
				for l < r {
					ans[l] = 'R'
					ans[r] = 'L'
					l++
					r--
				}
				lastR = -1
			} else {
				for j := lastRL + 1; j < i; j++ {
					ans[j] = 'L'
				}
			}
		}

		if d == 'R' {
			if lastR != -1 {
				for j := i - 1; j > lastR; j-- {
					ans[j] = 'R'
				}
			}
			lastR = i
		}
	}

	if lastR != -1 {
		for j := len(dominoes) - 1; j > lastR; j-- {
			ans[j] = 'R'
		}
	}

	return string(ans)
}
