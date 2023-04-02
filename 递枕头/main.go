package main

import "fmt"

func main() {
	fmt.Println(passThePillow(4, 5))
	fmt.Println(passThePillow(3, 2))
}

func passThePillow(n int, time int) int {
	ans := 1
	reverse := false
	for time != 0 {
		if ans == n {
			reverse = true
		} else if ans == 1 {
			reverse = false
		}
		switch reverse {
		case true:
			ans--
		case false:
			ans++
		}
		time--
		//fmt.Println(ans, time)
	}
	return ans
}
