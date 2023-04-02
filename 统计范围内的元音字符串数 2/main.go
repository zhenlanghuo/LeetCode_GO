package main

func main() {

}

func vowelStrings(words []string, left int, right int) int {

	m := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	ans := 0
	for i := left; i <= right; i++ {
		if m[words[i][0]] && m[words[i][len(words[i])-1]] {
			ans++
		}
	}

	return ans
}
