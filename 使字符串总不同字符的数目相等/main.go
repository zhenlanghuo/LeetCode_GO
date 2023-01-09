package main

import "fmt"

func main() {
	fmt.Println(isItPossible("ac", "b"))
	fmt.Println(isItPossible("abcc", "aab"))
	fmt.Println(isItPossible("abcde", "fghij"))
	fmt.Println(isItPossible("aa", "ab"))
}

func isItPossible(word1 string, word2 string) bool {
	word1Map := make(map[byte]int)
	word2Map := make(map[byte]int)

	for i := 0; i < 26; i++ {
		word1Map[byte('a')+byte(i)] = 0
		word2Map[byte('a')+byte(i)] = 0
	}

	for _, b := range word1 {
		word1Map[byte(b)]++
	}

	for _, b := range word2 {
		word2Map[byte(b)]++
	}

	for i := 0; i < 26; i++ {
		b1 := 'a' + byte(i)
		if word1Map[b1] == 0 {
			continue
		}
		for j := 0; j < 26; j++ {
			b2 := 'a' + byte(j)
			if word2Map[b2] == 0 {
				continue
			}

			word1Map[byte(b1)]--
			word2Map[byte(b2)]--
			word1Map[byte(b2)]++
			word2Map[byte(b1)]++
			if chars(word1Map) == chars(word2Map) {
				return true
			}
			word1Map[byte(b1)]++
			word2Map[byte(b2)]++
			word1Map[byte(b2)]--
			word2Map[byte(b1)]--
		}
	}

	return false
}

func chars(wordMap map[byte]int) int {
	result := 0
	for _, v := range wordMap {
		if v > 0 {
			result++
		}
	}
	return result
}
