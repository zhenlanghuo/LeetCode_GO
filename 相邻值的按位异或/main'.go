package main

import "fmt"

func main() {
	fmt.Println(doesValidArrayExist([]int{1, 1, 0}))
	//fmt.Println(doesValidArrayExist([]int{1, 0}))
}

func doesValidArrayExist(derived []int) bool {

	n := len(derived)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}

	ans[0] = 0
	for i, d := range derived {
		temp := 0
		if d == 0 {
			temp = ans[i]
		} else {
			temp = (ans[i] + 1) % 2
		}
		//fmt.Println(i, temp, ans)
		if ans[(i+1)%n] != -1 && temp != ans[(i+1)%n] {
			return false
		}
		ans[(i+1)%n] = temp
	}
	//fmt.Println(ans)

	return true
}
