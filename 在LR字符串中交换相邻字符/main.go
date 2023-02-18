package main

import "fmt"

func main() {
	fmt.Println(canTransform("RXXLRXRXL", "XRLXXRRLX"))
}

func canTransform(start string, end string) bool {
	if len(start) != len(end) {
		return false
	}

	n := len(start)
	i, j := 0, 0

	for i < n && j < n {
		for i < n && start[i] == 'X' {
			i++
		}

		for j < n && end[j] == 'X' {
			j++
		}

		if i >= n || j >= n {
			break
		}

		if start[i] != end[j] {
			return false
		}

		if start[i] == 'L' && i < j {
			return false
		}

		if start[i] == 'R' && i > j {
			return false
		}

		i++
		j++
	}

	fmt.Println(i, j)

	if i >= n {
		for j < n {
			if end[j] != 'X' {
				return false
			}
			j++
		}
	} else {
		for i < n {
			if start[i] != 'X' {
				return false
			}
			i++
		}
	}

	return true
}
