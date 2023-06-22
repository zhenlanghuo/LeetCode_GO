package main

func main() {

}

func isFascinating(n int) bool {

	m := make([]int, 10)

	for i := 1; i <= 3; i++ {
		tempN := n * i
		for tempN != 0 {
			m[tempN%10]++
			tempN = tempN / 10
		}
	}

	if m[0] != 0 {
		return false
	}

	for i := 1; i <= 9; i++ {
		if m[i] != 1 {
			return false
		}
	}

	return true
}
