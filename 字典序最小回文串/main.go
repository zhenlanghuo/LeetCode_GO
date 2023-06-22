package main

func main() {

}

func makeSmallestPalindrome(s string) string {
	bytes := []byte(s)

	l, r := 0, len(bytes)-1
	for l < r {
		if bytes[l] < bytes[r] {
			bytes[r] = bytes[l]
		} else {
			bytes[l] = bytes[r]
		}
		l++
		r--
	}

	return string(bytes)
}
