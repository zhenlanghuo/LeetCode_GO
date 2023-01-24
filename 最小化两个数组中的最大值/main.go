package main

import "fmt"

func main() {
	fmt.Println(minimizeSet(2, 7, 1, 3))
	fmt.Println(minimizeSet(3, 5, 2, 1))
	fmt.Println(minimizeSet(2, 4, 8, 2))
	fmt.Println(minimizeSet(95355, 5748, 652424220, 347575780))

}

func minimizeSet(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {
	lcm := divisor1 * divisor2 / gcd(divisor1, divisor2)

	check := func(x int) bool {
		return x-x/divisor1-x/divisor2+x/lcm >= max(uniqueCnt1-x/divisor2+x/lcm, 0)+max(uniqueCnt2-x/divisor1+x/lcm, 0)
	}

	left := 1
	right := (uniqueCnt1+uniqueCnt2)*2 + 1

	for left <= right {
		//fmt.Println(left, right)
		mid := (left + right) / 2
		if check(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func gcd(a, b int) int {
	for a > 0 {
		a, b = b%a, a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
