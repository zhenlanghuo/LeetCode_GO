package main

import "fmt"

func main() {
	fmt.Println(len("this is really a very awesome message"))
	fmt.Println(splitMessage("this is really a very awesome message", 8))
	fmt.Println(splitMessage("short message", 15))
}

func splitMessage(message string, limit int) []string {
	i, ca, tailLen := 1, 0, 0
	for ; ; i++ {
		if i < 10 {
			tailLen = 5
		} else if i < 100 {
			if i == 10 {
				ca -= 9
			}
			tailLen = 7
		} else if i < 1000 {
			if i == 100 {
				ca -= 99
			}
			tailLen = 9
		} else {
			if i == 1000 {
				ca -= 999
			}
			tailLen = 11
		}
		if limit-tailLen <= 0 {
			return nil
		}
		ca += limit - tailLen
		if ca < len(message) {
			continue
		}

		ans := make([]string, i)
		for j := range ans {
			tail := fmt.Sprintf("<%d/%d>", j+1, i)
			// fmt.Println(j, ans, message, tail)
			l := limit - len(tail)
			if j == len(ans)-1 {
				l = len(message)
			}
			ans[j] = fmt.Sprintf("%s%s", message[:l], tail)
			message = message[l:]
		}
		return ans
	}
}
