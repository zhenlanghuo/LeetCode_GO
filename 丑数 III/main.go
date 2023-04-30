package main

func main() {
	//fmt.Println(nthUglyNumber())
}

func nthUglyNumber(n int, a int, b int, c int) int {
	lcmAB := lcm(a, b)
	lcmAC := lcm(a, c)
	lcmBC := lcm(b, c)
	lcmABC := lcm(lcmAB, c)

	check := func(x int) bool {
		return x/a+x/b+x/c-x/lcmAC-x/lcmBC-x/lcmAB+x/lcmABC >= n
	}

	left, right := 0, int(2e9)
	for left <= right {
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
	for a != 0 {
		a, b = b%a, a
	}

	return b
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
