package main

import "fmt"

func main() {
	fmt.Println(distMoney(17, 2))
	fmt.Println(distMoney(1, 2))
	fmt.Println(distMoney(16, 2))
}

func distMoney(money int, children int) int {

	if money < children {
		return -1
	}

	ans := 0
	for {
		if children == 1 {
			if money == 4 {
				ans--
			}
			if money == 8 {
				ans++
			}
			break
		}

		if money-8 >= children-1 {
			ans++
			money -= 8
			children--
		} else {
			break
		}
	}

	return ans
}
