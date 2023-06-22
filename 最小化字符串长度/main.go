package main

func main() {

}

func minimizedStringLength(s string) int {
	m := make(map[byte]struct{})
	for i := 0; i < len(s); i++ {
		m[s[i]] = struct{}{}
	}

	return len(m)
}
