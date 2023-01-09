package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(replaceWords([]string{"cat","bat","rat"}, "the cattle was rattled by the battery"))
	fmt.Println(replaceWords([]string{"a","b","c"}, "aadsfasf absbs bbab cadsfafs"))
}

type trie map[byte]trie

func replaceWords(dictionary []string, sentence string) string {
	root := make(trie)
	for _, word := range dictionary {
		cur := root
		for _, b := range word {
			if cur[byte(b)] == nil {
				cur[byte(b)] = make(trie)
			}
			cur = cur[byte(b)]
		}
		cur['#'] = make(trie)
	}

	words := strings.Split(sentence, " ")
	for i, word := range words {
		cur := root
		for j, b := range word {
			if cur == nil {
				break
			}
			if cur['#'] != nil {
				words[i] = word[:j]
				break
			}

			cur = cur[byte(b)]
		}
	}

	return strings.Join(words, " ")
}