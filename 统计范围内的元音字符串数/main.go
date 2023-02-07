package main

func main() {

}

func vowelStrings(words []string, queries [][]int) []int {
	m := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

	count := make([]int, len(words)+1)
	for i, word := range words {
		count[i+1] = count[i]
		if m[word[0]] && m[word[len(word)-1]] {
			count[i+1]++
		}
	}

	result := make([]int, 0, len(queries))
	for _, query := range queries {
		result = append(result, count[query[1]+1]-count[query[0]])
	}

	return result
}
