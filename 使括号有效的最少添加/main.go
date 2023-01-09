package main

func main() {

}

func minAddToMakeValid(s string) int {

	leftCount := 0
	result := 0
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b == ')' {
			if leftCount == 0 {
				result++
			} else {
				leftCount--
			}
		}

		if b == '(' {
			leftCount++
		}
	}

	return result + leftCount
}
